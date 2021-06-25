package main

import (
	"go_demo/src/net_demo/dynamic/msg"
	"go_demo/src/net_demo/service/facade"
	"go_demo/src/net_demo/service/log"
	"go_demo/src/net_demo/service/net/tcp"
	"go_demo/src/net_demo/service/statistics"
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	log.InitLog("test", zapcore.DebugLevel)

	statisticsService := statistics.NewStatisticsService()
	statisticsService.Startup()

	serviceFacade := facade.NewServiceFacade()
	serviceFacade.RegisterService(statisticsService)
	//--------------------------------------------------
	var msgConfig = tcp.NewMsgConfig()
	var msgHandler = &tcp.ClientMsgHandler{}
	// 消息注册
	msgConfig.AddMsgRelation((int32)(msg.LoginDown_CommondId), &msg.LoginDown{}, msgHandler)

	msgConfig.AddMsgRelation((int32)(msg.LoginUp_CommondId), &msg.LoginUp{}, nil)
	//--------------------------------------------------
	var netServiceConfig = tcp.NewNetServiceConfig("net_demo/config/client.toml", msgConfig, serviceFacade)

	go func() {
		var n = 1
		var cnt int64 = 0
		for {
			//每隔1秒创建n个链接
			for k := 0; k < n; k++ {
				var tcpClient = tcp.NewTcpClient(netServiceConfig)
				tcpClient.Startup()

				var session = tcpClient.GetSession()
				cnt++

				msg := &msg.LoginUp{}
				session.Send(msg)
			}
			time.Sleep(time.Duration(1000) * time.Millisecond)
		}
	}()

	var i int
	var explain = "----------------------------\n" +
		"输入指令序号\n" +
		"1.退出测试\n" +
		"2.查看指令耗时"
	for {
		fmt.Println(explain)
		fmt.Scan(&i)
		var isOver = false
		switch i {
		case 1:
			isOver = true
			break
		case 2:
			statisticsService.Show()
			break
		default:
			break
		}
		if isOver {
			break
		}
	}
}
