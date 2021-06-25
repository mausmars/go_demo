package main

import (
	discovery "go_demo/src/zookeeper_demo"
	"fmt"
)

func main() {
	zk1 := discovery.ServiceDiscoveryZK{
		RootPath: "/game1",
	}
	zk1.Startup()
	zk2 := discovery.ServiceDiscoveryZK{
		RootPath: "/game2",
	}
	zk2.Startup()

	var i int
	var explain = "----------------------------\n" +
		"输入指令序号\n" +
		"0.退出测试 1.查看结点 2.注册节点 3.注销节点 4.修改数据\n"
	for {
		fmt.Println(explain)
		fmt.Scan(&i)
		var isOver = false
		switch i {
		case 0:
			zk1.Shutdown()
			zk2.Shutdown()
			isOver = true
			break
		case 1:
			zk1.PrintNodes()
			zk2.PrintNodes()
			break
		case 2:
			nodeData1 := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8001,
				ServiceName: "nodeData1_service1_1",
			}
			zk1.Register(nodeData1)

			nodeData2 := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8001,
				ServiceName: "nodeData2_service1_1",
			}
			zk2.Register(nodeData2)
			break
		case 3:
			zk1.Unregister("service1")
			zk2.Unregister("service1")
			break
		case 4:
			nodeData1 := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8002,
				ServiceName: "nodeData1_service1_2",
			}
			zk1.ModifyData(nodeData1)

			nodeData2 := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8002,
				ServiceName: "nodeData2_service1_2",
			}
			zk2.ModifyData(nodeData2)
			break
		}
		if isOver {
			break
		}
	}
}
