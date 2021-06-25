package net

import "net"

type NetServer struct {
	config       NetConfig
	childHandler []IChannelHandler
	listener     net.Listener
}

func (s *NetServer) Startup() {
	listener, err := net.Listen(s.config.netProtocol, s.config.host)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	s.accept()

}

func (s *NetServer) Shutdown() {

}

func (s *NetServer) accept() {
	for i := 0; i < s.config.acceptCoroutineNum; i++ {
		go func() {
			conn, err := s.listener.Accept()
			if err != nil {
				panic(err)
			}
			//tcp设置
			s.setOption(conn)

		}()
	}
}

func (s *NetServer) setOption(conn net.Conn) {
	c := conn.(*net.TCPConn)
	c.SetNoDelay(s.config.nodely)
	c.SetLinger(s.config.linger)
	c.SetKeepAlive(s.config.keepalive)

	//c.SetReadDeadline()
	//c.SetWriteDeadline()
	//c.SetDeadline()

	//c.SetReadBuffer()
	//c.SetWriteBuffer()
}
