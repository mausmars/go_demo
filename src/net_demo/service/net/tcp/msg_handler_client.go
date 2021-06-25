package tcp

import (
	"go_demo/src/net_demo/dynamic/msg"
	"go_demo/src/net_demo/service/statistics"
	"go_demo/src/net_demo/service/util"
	"fmt"
	"time"
)

//默认处理器
type ClientMsgHandler struct {
}

func (h *ClientMsgHandler) handle(session ISession, m interface{}) {
	util.Try(func() {
		msgName := util.GetMsgName(m)
		switch msgName {
		case "LoginDown":
			//记录发送耗时
			ci := h.getCommandInfo(session)
			if ci != nil {
				if ci.Command == "LoginUp" {
					h.recorder(session)
				}
			}
			tm := m.(*msg.LoginDown)
			fmt.Println("handle msg=", tm)
			break
		default:
			fmt.Println("handle default msg=", m)
			break
		}
	}, func(e interface{}) {
		fmt.Println("Client handler Error!!! ", e)
	})
}

func (h *ClientMsgHandler) getCommandInfo(session ISession) *statistics.CommandInfo {
	v := session.GetAttach(SessionAttach_Queue)
	if v == nil {
		return nil
	}
	commandQueue := v.(*DoubleLinkedQueue)
	node := commandQueue.Peek()
	if v == nil {
		return nil
	}
	return node.data.(*statistics.CommandInfo)
}

func (h *ClientMsgHandler) recorder(session ISession) {
	v := session.GetAttach(SessionAttach_Queue)
	if v == nil {
		return
	}

	statisticsService := session.GetNetServiceConfig().serviceFacade.GetStatisticsService()
	if statisticsService == nil {
		return
	}

	commandQueue := v.(*DoubleLinkedQueue)
	node := commandQueue.Pop()

	var ci = node.data.(*statistics.CommandInfo)
	ci.IsSuccess = true
	ci.Time = time.Now().Sub(ci.SendTime).Milliseconds()

	statisticsService.Recorder(ci)
}
