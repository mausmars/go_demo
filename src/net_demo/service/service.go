package service

const (
	ServiceType_Queue      = "Queue"
	ServiceType_Statistics = "Statistics"
	ServiceType_NetServer  = "NetServer"
	ServiceType_NetClient  = "NetClient"
	ServiceType_FlumeLog   = "FlumeLog"

	ServiceDefaultName = "Default"
)

type IService interface {
	Startup()
	Shutdown()

	ServiceType() string
	ServiceName() string
}

/**
服务管理
*/
type ServiceContext struct {
	//{服务类型:{名字:服务}}
	serviceMap map[string]map[string]IService
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		serviceMap: make(map[string]map[string]IService),
	}
}

func (p *ServiceContext) GetTypeService(serviceType string) map[string]IService {
	return p.serviceMap[serviceType]
}

func (p *ServiceContext) GetDefaultService(serviceType string) IService {
	m, ok := p.serviceMap[serviceType]
	if !ok {
		return nil
	}
	return m[ServiceDefaultName]
}

func (p *ServiceContext) GetServiceByName(serviceType string, serviceName string) IService {
	m, ok := p.serviceMap[serviceType]
	if !ok {
		return nil
	}
	return m[serviceName]
}

func (p *ServiceContext) RegisterService(service IService) {
	m, ok := p.serviceMap[service.ServiceType()]
	if !ok {
		m = make(map[string]IService)
		p.serviceMap[service.ServiceType()] = m
	}
	m[service.ServiceName()] = service
}

//--------------------------------------------------------
//服務端网络服务
type INetServerService interface {
	IService
}

//网络客户端服务
type INetClientService interface {
	IService
}

//日志服务
type ILogService interface {
	IService

	IsDebug() bool
	IsInfo() bool

	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})

	GetRawLogger() interface{}
}
