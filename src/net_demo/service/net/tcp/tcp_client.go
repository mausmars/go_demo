package tcp

import (
	"go_demo/src/net_demo/service"
	"go_demo/src/net_demo/service/util"
	"fmt"
	"net"
	"time"

	//"github.com/gogo/protobuf/proto"
)

type TcpClient struct {
	serviceName string
	conf        *NetServiceConfig
	session     ISession
}

func NewTcpClient(conf *NetServiceConfig) *TcpClient {
	var c = &TcpClient{
		serviceName: service.ServiceDefaultName,
		conf:        conf,
	}
	return c
}

func (c *TcpClient) Startup() {
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", c.conf.netConfig.Host); err != nil; conn, err = net.Dial("tcp", c.conf.netConfig.Host) {
		fmt.Println("connect", c.conf.netConfig.Host, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect", c.conf.netConfig.Host, "success")

	//tcp设置
	c.setOption(conn)
	var session = NewClientSession(conn, c.conf)
	c.SetSession(session)
	//协程读取处理消息
	go c.readMessage(session)
}

func (s *TcpClient) ServiceType() string {
	return service.ServiceType_NetClient
}

func (s *TcpClient) ServiceName() string {
	return s.serviceName
}

func (c *TcpClient) SetSession(session ISession) {
	c.session = session
}
func (c *TcpClient) GetSession() ISession {
	return c.session
}

func (s *TcpClient) setOption(conn net.Conn) {
	c := conn.(*net.TCPConn)
	c.SetNoDelay(s.conf.netConfig.Nodely)
	c.SetLinger(s.conf.netConfig.Linger)
	c.SetKeepAlive(s.conf.netConfig.Keepalive)

	//c.SetReadDeadline()
	//c.SetWriteDeadline()

	//c.SetDeadline()
	//c.SetReadBuffer()
	//c.SetWriteBuffer()
}

//接收消息
func (s *TcpClient) readMessage(session ISession) {
	util.Try(func() {
		var netDecoder = NewNetDecoder(session.GetNetServiceConfig().netConfig.BuffSize)
		for {
			cnt, err := session.Read(netDecoder.buf)
			if err != nil {
				panic(err)
			}
			pData := netDecoder.buf[:cnt]
			//fmt.Println("read ", pData)
			netDecoder.buffer.Write(pData)

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
		fmt.Println("Client net handle Error!!! ", e)
		session.Close()
		//关闭session 处理
		s.close(session)
	})
}

func (s *TcpClient) close(session ISession) {
	session.SetAttach("SessionClose", true)
}

//func StartClient() {
//	strIP := "localhost:12345"
//	var conn net.Conn
//	var err error
//
//	//连接服务器
//	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
//		fmt.Println("connect", strIP, "fail")
//		time.Sleep(time.Second)
//		fmt.Println("reconnect...")
//	}
//	fmt.Println("connect", strIP, "success")
//	defer conn.Close()
//
//	//发送消息
//	var cnt int64 = 0
//
//	for i := 1; i < 20; i++ {
//		cnt++
//
//		stSend := &queueproto.PushUp{
//			Member: *proto.Int64(cnt),
//		}
//		//protobuf编码
//		pData, err := proto.Marshal(stSend)
//		if err != nil {
//			panic(err)
//		}
//
//		var netProto = new(NetProto)
//		netProto.SetMsgId(int32(queueproto.PushUp_CommondId))
//		netProto.SetBody(pData)
//
//		var buffer = bytes.NewBuffer([]byte{})
//		netProto.write(buffer)
//
//		//发送
//		conn.Write(buffer.Bytes())
//
//		//if (i%10 == 0) {
//		//	time.Sleep(50 * time.Millisecond)
//		//}
//
//		fmt.Println("send message!", buffer.Bytes())
//	}
//	time.Sleep(time.Duration(2) * time.Minute)
//}
