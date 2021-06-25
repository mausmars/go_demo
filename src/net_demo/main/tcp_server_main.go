package main

import (
	"go_demo/src/net_demo/dynamic/msg"
	"go_demo/src/net_demo/service/facade"
	"go_demo/src/net_demo/service/log"
	"go_demo/src/net_demo/service/net/tcp"
	"fmt"
	"go.uber.org/zap/zapcore"
)

func main() {
	log.InitLog("test", zapcore.DebugLevel)

	serviceFacade := facade.NewServiceFacade()
	//-----------------------------------------
	var msgConfig = tcp.NewMsgConfig()
	var msgHandler = &tcp.ServerMsgHandler{}
	// 消息注册
	msgConfig.AddMsgRelation((int32)(msg.LoginUp_CommondId), &msg.LoginUp{}, msgHandler)

	msgConfig.AddMsgRelation((int32)(msg.LoginDown_CommondId), &msg.LoginDown{}, nil)
	//-----------------------------------------
	var netServiceConfig = tcp.NewNetServiceConfig("net_demo/config/server.toml", msgConfig, serviceFacade)
	var tcpServer = tcp.NewTcpServer(netServiceConfig)
	go tcpServer.Startup()

	var i int
	var explain = "----------------------------\n" +
		"输入指令序号\n" +
		"1.退出测试"
	for {
		fmt.Println(explain)
		fmt.Scan(&i)
		var isOver = false
		switch i {
		case 1:
			tcpServer.Shutdown()
			isOver = true
			break
		default:
			break
		}
		if isOver {
			break
		}
	}
}
