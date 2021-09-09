package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
)

const (
	DefaultHeadLength  = 6
	TestCommandId_Up   = 1
	TestCommandId_Down = 2
)

//impl ICodec
type NetDataProtocol struct {
	CommandId  uint16 //2
	DataLength uint32 //4
	Data       []byte
}

// Encode ...
func (cc *NetDataProtocol) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	result := make([]byte, 0)
	buffer := bytes.NewBuffer(result)
	// take out the param
	item := c.Context().(NetDataProtocol)
	if err := binary.Write(buffer, binary.BigEndian, item.CommandId); err != nil {
		s := fmt.Sprintf("Pack type error , %v", err)
		return nil, errors.New(s)
	}
	dataLen := uint32(len(buf))
	if err := binary.Write(buffer, binary.BigEndian, dataLen); err != nil {
		s := fmt.Sprintf("Pack datalength error , %v", err)
		return nil, errors.New(s)
	}
	if dataLen > 0 {
		if err := binary.Write(buffer, binary.BigEndian, buf); err != nil {
			s := fmt.Sprintf("Pack data error , %v", err)
			return nil, errors.New(s)
		}
	}
	return buffer.Bytes(), nil
}

// Decode ...
func (cc *NetDataProtocol) Decode(c gnet.Conn) ([]byte, error) {
	// parse header
	headerLen := DefaultHeadLength // uint16+uint16+uint32
	if size, header := c.ReadN(headerLen); size == headerLen {
		byteBuffer := bytes.NewBuffer(header)
		var commandId uint16
		var dataLength uint32
		_ = binary.Read(byteBuffer, binary.BigEndian, &commandId)
		_ = binary.Read(byteBuffer, binary.BigEndian, &dataLength)
		dataLen := int(dataLength) // max int32 can contain 210MB payload
		protocolLen := headerLen + dataLen
		if dataSize, data := c.ReadN(protocolLen); dataSize == protocolLen {
			c.ShiftN(protocolLen)
			return data[headerLen:], nil
		}
		return nil, errors.New("not enough payload data")
	}
	return nil, errors.New("not enough header data")
}
