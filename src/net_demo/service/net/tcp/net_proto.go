package tcp

import (
	"bytes"
	"encoding/binary"
)

type ReadState int

/**
网络协议
 */
type NetProto struct {
	length int32
	msgId  int32
	body   []byte
}

func (p *NetProto) SetMsgId(msgId int32) {
	p.msgId = msgId
}

func (p *NetProto) GetMsgId() int32 {
	return p.msgId
}

func (p *NetProto) SetBody(body []byte) {
	p.body = body
}

func (p *NetProto) GetBody() []byte {
	return p.body
}

func (p *NetProto) write(buffer *bytes.Buffer) {
	p.length = int32(8 + len(p.body)) //8字节的长度
	binary.Write(buffer, binary.BigEndian, p.length)
	binary.Write(buffer, binary.BigEndian, p.msgId)
	buffer.Write(p.body)
}

func (p *NetProto) readLen(buffer *bytes.Buffer) {
	binary.Read(buffer, binary.BigEndian, &p.length)
}

func (p *NetProto) readBody(buffer *bytes.Buffer) {
	binary.Read(buffer, binary.BigEndian, &p.msgId)
	p.body = make([]byte, p.length-8)
	binary.Read(buffer, binary.BigEndian, &p.body)
}
