package main

import (
	discovery "go_demo/src/zookeeper_demo"
	"fmt"
)

func main() {
	zk := discovery.ServiceDiscoveryZK{}
	zk.Startup()

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
			zk.Shutdown()
			isOver = true
			break
		case 1:
			zk.PrintNodes()
			break
		case 2:
			nodeData := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8001,
				ServiceName: "test1",
			}
			nodeData2 := &discovery.NodeData{
				NodeName:    "service2",
				Host:        "127.0.0.1",
				Port:        8001,
				ServiceName: "test2",
			}
			zk.Register(nodeData)
			zk.Register(nodeData2)
			break
		case 3:
			zk.Unregister("service1")
			zk.Unregister("service2")
			break
		case 4:
			nodeData := &discovery.NodeData{
				NodeName:    "service1",
				Host:        "127.0.0.1",
				Port:        8002,
				ServiceName: "test2",
			}
			nodeData2 := &discovery.NodeData{
				NodeName:    "service2",
				Host:        "127.0.0.1",
				Port:        8002,
				ServiceName: "test2",
			}
			zk.ModifyData(nodeData)
			zk.ModifyData(nodeData2)
			break
		}
		if isOver {
			break
		}
	}
}
