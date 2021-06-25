package tcp

import (
	"go_demo/src/net_demo/dynamic/msg"
	"go_demo/src/net_demo/service/util"
	"fmt"
)

//默认处理器
type ServerMsgHandler struct {
}

func (h *ServerMsgHandler) handle(session ISession, m interface{}) {
	util.Try(func() {
		msgName := util.GetMsgName(m)
		switch msgName {
		case "LoginUp":
			tm := m.(*msg.LoginUp)
			fmt.Println("handle msg=", tm)
			//session.SetAttach("member", tm.Member)
			//发送回复
			resp := &msg.LoginDown{}
			session.Send(resp)
			break
		default:
			fmt.Println("handle default msg=", m)
			break
		}
	}, func(e interface{}) {
		fmt.Println("Server handlere Error!!! ", e)
	})
}
