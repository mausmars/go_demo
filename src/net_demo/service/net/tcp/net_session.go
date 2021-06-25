package tcp

import (
	"bytes"
	"go_demo/src/net_demo/service/log"
	"go_demo/src/net_demo/service/statistics"
	"go_demo/src/net_demo/service/util"
	"github.com/golang/protobuf/proto"
	"net"
	"time"
)

const (
	SessionAttach_Queue = "Queue"
)

type ISession interface {
	Send(msg interface{})
	GetNetServiceConfig() *NetServiceConfig

	Close()
	Read(b []byte) (n int, err error)

	GetAttach(k string) interface{}
	SetAttach(k string, v interface{})
}

type Session struct {
	conn   net.Conn
	conf   *NetServiceConfig
	attach map[string]interface{}
}

func NewClientSession(conn net.Conn, conf *NetServiceConfig) ISession {
	session := &Session{
		conn:   conn,
		conf:   conf,
		attach: make(map[string]interface{}),
	}
	session.attach[SessionAttach_Queue] = new(DoubleLinkedQueue)
	return session
}

func NewServerSession(conn net.Conn, conf *NetServiceConfig) ISession {
	session := &Session{
		conn:   conn,
		conf:   conf,
		attach: make(map[string]interface{}),
	}
	return session
}
func (s *Session) GetAttach(k string) interface{} {
	return s.attach[k]
}
func (s *Session) SetAttach(k string, v interface{}) {
	s.attach[k] = v
}

//为了统计用
func (s *Session) statistics(msg interface{}) {
	v, ok := s.attach[SessionAttach_Queue]
	if !ok {
		return
	}
	msgName := util.GetMsgName(msg)
	commandInfo := &statistics.CommandInfo{
		Command:  msgName,
		SendTime: time.Now(),
	}
	commandQueue := v.(*DoubleLinkedQueue)
	commandQueue.Push(commandInfo)
}

func (s *Session) Send(msg interface{}) {
	s.statistics(msg) //为了统计用

	bodyData, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		panic(err)
	}
	msgRelation := s.conf.msgConfig.getMsgRelationByType(msg)
	if msgRelation == nil {
		//消息不存在映射
		msgName := util.GetMsgName(msg)
		log.Logger.Error("Msg don't have map! ", msgName)
		return
	}
	var netProto = new(NetProto)
	netProto.SetMsgId(msgRelation.cid)
	netProto.SetBody(bodyData)

	var buffer = bytes.NewBuffer([]byte{})
	netProto.write(buffer)
	//发送CommandInfo
	s.conn.Write(buffer.Bytes())
}

func (s *Session) GetNetServiceConfig() *NetServiceConfig {
	return s.conf
}

func (s *Session) Close() {
	s.conn.Close()
}

func (s *Session) Read(b []byte) (n int, err error) {
	return s.conn.Read(b)
}
