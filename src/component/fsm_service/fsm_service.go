package fsm_service

import (
	"fmt"
	"github.com/looplab/fsm"
)

type FSMService struct {
	
}

func NewFSMService() *FSMService {
	return &FSMService{}
}

func (s *FSMService) Startup() {
	fsm := fsm.NewFSM(
		"closed", // 初始状态
		fsm.Events{
			// 定义事件，事件名称，状态src，状态dst
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			// 事件回调
			"before_event": func(event *fsm.Event) {
				fmt.Println("我要开/关门")
			},
			"before_open": func(event *fsm.Event) {
				fmt.Println("门要开了")
			},
			"before_close": func(event *fsm.Event) {
				fmt.Println("门要关了")
			},
			"leave_state": func(event *fsm.Event) {
				fmt.Println("结束了")
			},
		},
	)
	//当前状态
	fmt.Println("当前状态",fsm.Current())
	//改变状态
	err := fsm.Event("open")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("当前状态",fsm.Current())
	//改变状态
	err = fsm.Event("close")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("当前状态",fsm.Current())
}


func (s *FSMService) Shutdown() {

}