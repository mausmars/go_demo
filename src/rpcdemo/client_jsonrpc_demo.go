package main

import (
	"go_demo/src/rpcdemo/rpcmodel"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	//连接远程rpc服务
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8092")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	for i:=0;i<10;i++{
		err2 := rpc.Call("Rect.Area", rpcmodel.Params{30, 100}, &ret)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(ret)
		err3 := rpc.Call("Rect.Perimeter", rpcmodel.Params{50, 100}, &ret)
		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Println(ret)
		time.Sleep(3*time.Second)
	}


}