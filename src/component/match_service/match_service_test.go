package match_service

import (
	"bytes"
	"container/list"
	"go_demo/src/net_demo/service/util"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMatchService(t *testing.T) {
	//匹配上回调
	matchService := NewMatchService(10, &MatchAttach{Rid: 1}, overMatchCallback)
	matchService.Startup()

	size := int64(25)

	go func() {
		for i := int64(1); i <= size; i++ {
			n := &MatchNode{
				NodeId:   i,
				Callback: matchStateCallback,
			}
			matchService.InsertNode(n)
			if i < 10 {
				time.Sleep(50 * time.Millisecond)
			} else {
				time.Sleep(500 * time.Millisecond)
			}

		}
	}()

	go func() {
		for {
			n := rand.Intn(int(size))
			matchService.RemoveNode(int64(n))
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			matchService.ShowNode()
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(25 * time.Second)
	matchService.Shutdown()
	time.Sleep(40 * time.Second)
	fmt.Println("over!!!")
}

func overMatchCallback(finishedType int, linkedlist *list.List, attach interface{}) {
	matchAttach := attach.(*MatchAttach)
	var buffer bytes.Buffer
	n := linkedlist.Front()
	for {
		if n == nil {
			break
		}
		matchNode := n.Value.(*MatchNode)
		buffer.WriteString(util.Int642String(matchNode.NodeId))
		buffer.WriteString(",")
		n = n.Next()
	}
	fmt.Println("匹配上回调 rid=", matchAttach.Rid, " finishedType=", finishedType, " 匹配上{", buffer.String(), "}")
}

func matchStateCallback(nodeId int64, state int) {
	switch state {
	case Insert:
		//成功进入匹配池
		fmt.Println("Insert", nodeId)
		break
	case Remove:
		//成功从匹配池移除，推送gs服通知
		fmt.Println("Remove", nodeId)
		break
	case Match:
		//匹配成功
		fmt.Println("Match", nodeId)
		break
	}
}
