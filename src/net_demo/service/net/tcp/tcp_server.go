package tcp

import (
	"go_demo/src/net_demo/service"
	"go_demo/src/net_demo/service/util"
	"fmt"
	"net"
	"time"
)

type TcpServer struct {
	serviceName string
	conf        *NetServiceConfig
	listener    net.Listener

	channelHandler IChannelHandler
}

func NewTcpServer(conf *NetServiceConfig) *TcpServer {
	return &TcpServer{
		serviceName: service.ServiceDefaultName,
		conf:        conf,
	}
}

func (s *TcpServer) Startup() {
	//监听
	listener, err := net.Listen("tcp", s.conf.netConfig.Host)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	fmt.Println("TcpServer start finished! ", s.conf.netConfig.Host)

	s.accept()
}

func (s *TcpServer) Shutdown() {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *TcpServer) ServiceType() string {
	return service.ServiceType_NetServer
}

func (s *TcpServer) ServiceName() string {
	return s.serviceName
}

func (s *TcpServer) accept() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		//tcp设置
		s.setOption(conn)

		var session = NewServerSession(conn, s.conf)
		//协程读取处理消息
		go s.handle(session)
	}
}

func (s *TcpServer) setOption(conn net.Conn) {
	c := conn.(*net.TCPConn)
	c.SetNoDelay(s.conf.netConfig.Nodely)
	c.SetLinger(s.conf.netConfig.Linger)
	c.SetKeepAlive(s.conf.netConfig.Keepalive)

	//c.SetReadDeadline(time.Now().Add(time.Duration(s.conf.netConfig.ReadTimeout) * time.Millisecond))
	//c.SetWriteDeadline()
	//c.SetDeadline()

	//c.SetReadBuffer()
	//c.SetWriteBuffer()
}

//接收消息
func (s *TcpServer) handle(session ISession) {
	util.Try(func() {
		var netDecoder = NewNetDecoder(session.GetNetServiceConfig().netConfig.BuffSize)
		for {
			//这里避免多次写入
			cnt, err := session.Read(netDecoder.buf)
			if err != nil {
				panic(err)
			}
			pData := netDecoder.buf[:cnt]
			//fmt.Println("read ", pData)
			netDecoder.buffer.Write(pData)

			//检测每次Client是否有数据传来
			for {
				//读消息
				state, msg := netDecoder.Decode(session)
				if !state {
					//fmt.Println("不能解析 ", msg)
					break
				} else {
					//处理业务
					//fmt.Println("需要处理业务 ", msg)
					conf := session.GetNetServiceConfig()
					msgRelation := conf.msgConfig.getMsgRelationByType(msg)
					// 消息处理
					msgRelation.handler.handle(session, msg)
				}
			}
		}
	}, func(e interface{}) {
		fmt.Println("Server net handle Error!!! ", e)
		session.Close()
		//关闭session 处理
		s.close(session)
	})
}

//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func heartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case fk := <-readerChannel:
		fmt.Println(conn.RemoteAddr().String(), "receive data string:", string(fk))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		//conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
		break
	case <-time.After(time.Second * 5):
		fmt.Println("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func (s *TcpServer) close(session ISession) {
	util.Try(func() {
		var member = session.GetAttach("member").(string)
		if member != "" {
			//queueService := s.conf.serviceContext.getService(ServiceType_QueueService)
			//queueService.(RedisQueueService).Remove(member)
		}
	}, func(e interface{}) {
		fmt.Println("Server close Error!!! ", e)
	})
}
