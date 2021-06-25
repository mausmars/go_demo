package facade

import (
	"go_demo/src/net_demo/service"
	"go_demo/src/net_demo/service/statistics"
)

type ServiceFacade struct {
	serviceContext *service.ServiceContext
}

func NewServiceFacade() *ServiceFacade {
	serviceContext := service.NewServiceContext()
	facade := &ServiceFacade{serviceContext: serviceContext}
	return facade
}

func (f *ServiceFacade) GetStatisticsService() *statistics.StatisticsService {
	s := f.serviceContext.GetDefaultService(service.ServiceType_Statistics)
	return s.(*statistics.StatisticsService)
}

func (f *ServiceFacade) GetNetClientService() service.INetClientService {
	s := f.serviceContext.GetDefaultService(service.ServiceType_NetClient)
	return s.(service.INetClientService)
}

func (f *ServiceFacade) GetNetServerService() service.INetServerService {
	s := f.serviceContext.GetDefaultService(service.ServiceType_NetServer)
	return s.(service.INetServerService)
}

func (f *ServiceFacade) RegisterService(service service.IService) {
	f.serviceContext.RegisterService(service)
}

func (f *ServiceFacade) GetServiceContext() *service.ServiceContext {
	return f.serviceContext
}
