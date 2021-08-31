package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"go_demo/src/gnet_demo/my_demo/protocol"
	"log"
	"sync"
	"time"
)

type TcpServer struct {
	*gnet.EventServer

	connectedSockets sync.Map
	async            bool
	tick             time.Duration
	workerPool       *goroutine.Pool
}

func (s *TcpServer) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	log.Printf("服务监听 %s (multi-cores: %t, loops: %d), "+
		"pushing data every %s ...\n", svr.Addr.String(), svr.Multicore, svr.NumEventLoop, s.tick.String())
	return
}

func (s *TcpServer) OnShutdown(svr gnet.Server) {
	log.Printf("关闭 addr= %s \n", svr.Addr.String())
}

func (s *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("开启一个连接 addr= %s \n", c.RemoteAddr().String())
	s.connectedSockets.Store(c.RemoteAddr().String(), c)
	return
}

func (s *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("关闭一个连接 addr= %s \n", c.RemoteAddr().String())
	s.connectedSockets.Delete(c.RemoteAddr().String())
	return
}

func (s *TcpServer) PreWrite() {
	log.Printf("准备写 \n")
}

func (s *TcpServer) Tick() (delay time.Duration, action gnet.Action) {
	log.Println("tick")
	s.connectedSockets.Range(func(key, value interface{}) bool {
		addr := key.(string)
		c := value.(gnet.Conn)
		c.AsyncWrite([]byte(fmt.Sprintf("heart beating to %s\n", addr)))
		return true
	})
	delay = s.tick
	return
}

func (s *TcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("frame:", string(frame))
	// store customize protocol header param using `c.SetContext()`
	item := protocol.NetDataProtocol{Version: protocol.DefaultProtocolVersion, ActionType: protocol.ActionData}
	c.SetContext(item)
	if s.async {
		data := append([]byte{}, frame...)
		_ = s.workerPool.Submit(func() {
			c.AsyncWrite(data)
		})
		return
	}
	out = frame
	return
}

func main() {
	port := 9000
	multicore := true
	interval := 1 * time.Second
	ticker := true
	async := true

	server := &TcpServer{
		async:      async,
		tick:       interval,
		workerPool: goroutine.Default(),
	}
	addr := fmt.Sprintf("tcp://:%d", port)
	log.Fatal(gnet.Serve(server, addr,
		gnet.WithMulticore(multicore),
		gnet.WithTicker(ticker),
		gnet.WithTCPKeepAlive(time.Minute*5),
		gnet.WithCodec(&protocol.NetDataProtocol{}),
	))
}
