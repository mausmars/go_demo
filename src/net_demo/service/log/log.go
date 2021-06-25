package log

import (
	"go_demo/src/net_demo/service"
	"fmt"
	"github.com/gemalto/flume"
	"go.uber.org/zap/zapcore"
)

// log设置为全局变量
var Logger service.ILogService

func InitLog(logName string, logLevel zapcore.Level) {
	s := NewFlumeLogService(logName, logLevel)
	s.Startup()
	Logger = s
}

type FlumeLogService struct {
	logger flume.Logger

	logName  string
	encoding string
	logLevel zapcore.Level
}

func NewFlumeLogService(logName string, logLevel zapcore.Level) *FlumeLogService {
	switch logLevel {
	case zapcore.DebugLevel, zapcore.InfoLevel, zapcore.ErrorLevel:
		break
	default:
		logLevel = zapcore.DebugLevel
		break
	}
	s := &FlumeLogService{
		logName:  logName,
		encoding: "ltsv",
		logLevel: logLevel,
	}
	return s
}

func (s *FlumeLogService) Startup() {
	err := flume.Configure(flume.Config{
		Development:  true,
		DefaultLevel: flume.Level(s.logLevel),
		Encoding:     s.encoding,
	})
	if err != nil {
		fmt.Println("logging config failed:", err.Error())
		return
	}
	s.logger = flume.New(s.logName)
}

func (s *FlumeLogService) Shutdown() {
}

func (s *FlumeLogService) ServiceType() string {
	return service.ServiceType_FlumeLog
}

func (s *FlumeLogService) ServiceName() string {
	return service.ServiceDefaultName
}

func (s *FlumeLogService) IsDebug() bool {
	return s.logger.IsDebug()
}

func (s *FlumeLogService) IsInfo() bool {
	return s.logger.IsInfo()
}

func (s *FlumeLogService) Debug(msg string, args ...interface{}) {
	s.logger.Debug(msg, args)
}

func (s *FlumeLogService) Info(msg string, args ...interface{}) {
	s.logger.Info(msg, args)
}

func (s *FlumeLogService) Error(msg string, args ...interface{}) {
	s.logger.Error(msg, args)
}

func (s *FlumeLogService) GetRawLogger() interface{} {
	return s.logger
}
