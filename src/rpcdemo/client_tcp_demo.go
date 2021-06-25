
package main

import (
	"go_demo/src/rpcdemo/rpcmodel"
	"net/rpc"
	"fmt"
	"log"
	"time"
)

func main() {
	//连接远程rpc服务
	//这里使用Dial，http方式使用DialHTTP，其他代码都一样
	rpc, err := rpc.Dial("tcp", "127.0.0.1:8091")
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