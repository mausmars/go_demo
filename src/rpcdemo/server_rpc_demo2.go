package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	"go_demo/src/rpcdemo/rpcmodel"
)

//go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
//go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码

func main() {
	rect := new(rpcmodel.Rect)
	//注册一个rect服务
	rpc.Register(rect)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		//把服务处理绑定到http协议上
		rpc.HandleHTTP()
		err := http.ListenAndServe(":8090", nil)
		wg.Wait()
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
	}()
	log.Println("http rpc service start success addr:8090")

	go func() {
		tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8091")
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
		for {
			conn, err3 := tcplisten.Accept()
			if err3 != nil {
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()
	log.Println("tcp rpc service start success addr:8091")

	go func() {
		tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8092")
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
		for {
			conn, err3 := tcplisten.Accept()
			if err3 != nil {
				continue
			}
			go jsonrpc.ServeConn(conn)
		}
	}()
	log.Println("tcp json-rpc service start success addr:8092")
	wg.Wait()
}