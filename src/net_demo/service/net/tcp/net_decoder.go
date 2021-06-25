package tcp

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"reflect"
)

const (
	ReadState_New = iota
	ReadState_Head
)

type NetDecoder struct {
	netProto *NetProto
	buffer   *bytes.Buffer
	buf      []byte
	state    int32
}

func NewNetDecoder(buffSize int32) *NetDecoder {
	var d = &NetDecoder{}
	d.buf = make([]byte, buffSize)
	d.buffer = bytes.NewBuffer(make([]byte, 0, buffSize))
	d.state = ReadState_New
	return d
}

func (d *NetDecoder) Decode(session ISession) (bool, interface{}) {
	if (d.netProto == nil) {
		d.netProto = new(NetProto)
	}

	if (d.state == ReadState_New) {
		if (d.buffer.Len() < 4) {
			d.state = ReadState_New
			return false, nil
		}
		d.netProto.readLen(d.buffer)
		if ((int32)(d.buffer.Len()) < d.netProto.length-4) {
			d.state = ReadState_Head
			return false, nil
		}
		d.netProto.readBody(d.buffer)

		var msgRelation = session.GetNetServiceConfig().msgConfig.commandMap[d.netProto.msgId]
		var stReceive = reflect.New(msgRelation.msgType.Elem()).Interface() //反射创建消息
		//protobuf解码
		err := proto.Unmarshal(d.netProto.body, stReceive.(proto.Message))
		if err != nil {
			panic(err)
		}

		d.state = ReadState_New
		d.netProto = nil

		return true, stReceive
	} else {
		if ((int32)(d.buffer.Len()) < d.netProto.length-4) {
			d.state = ReadState_Head
			return false, nil
		}
		d.netProto.readBody(d.buffer)

		var conf = session.GetNetServiceConfig()
		var msgRelation = conf.msgConfig.getMsgRelationByMsgId(d.netProto.msgId)
		var stReceive = reflect.New(msgRelation.msgType.Elem()).Interface() //反射创建消息
		//protobuf解码
		err := proto.Unmarshal(d.netProto.body, stReceive.(proto.Message))
		if err != nil {
			panic(err)
		}
		d.state = ReadState_New
		d.netProto = nil

		return true, stReceive
	}
}
