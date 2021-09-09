package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go_demo/src/gnet_demo/my_demo/protocol"
	"io"
	"log"
	"net"
)

type TcpClient struct {
	host     string
	port     uint
	protocol *protocol.NetDataProtocol

	conn net.Conn
}

func (s *TcpClient) Startup() {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	s.conn = conn
	go func() {
		for {
			response, err := ClientDecode(conn)
			if err != nil {
				log.Printf("ClientDecode error, %v\n", err)
			}
			log.Printf("receive, msg=%s\n", string(response.Data))
		}
	}()
}

func (s *TcpClient) Shutdown() {
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *TcpClient) Send(data []byte) {
	pbdata, err := ClientEncode(protocol.TestCommandId_Up, data)
	if err != nil {
		panic(err)
	}
	s.conn.Write(pbdata)
}

// Example command: go run client.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go func() {
		for {
			response, err := ClientDecode(conn)
			if err != nil {
				log.Printf("ClientDecode error, %v\n", err)
			}
			log.Printf("receive, msg=%s\n", string(response.Data))
		}
	}()
	data := []byte("hello")
	pbdata, err := ClientEncode(protocol.TestCommandId_Up, data)
	if err != nil {
		panic(err)
	}
	conn.Write(pbdata)
	data = []byte("world")
	pbdata, err = ClientEncode(protocol.TestCommandId_Up, data)
	if err != nil {
		panic(err)
	}
	conn.Write(pbdata)
	select {}
}

// ClientEncode :
func ClientEncode(actionType uint16, data []byte) ([]byte, error) {
	result := make([]byte, 0)
	buffer := bytes.NewBuffer(result)
	if err := binary.Write(buffer, binary.BigEndian, actionType); err != nil {
		s := fmt.Sprintf("Pack type error , %v", err)
		return nil, errors.New(s)
	}
	dataLen := uint32(len(data))
	if err := binary.Write(buffer, binary.BigEndian, dataLen); err != nil {
		s := fmt.Sprintf("Pack datalength error , %v", err)
		return nil, errors.New(s)
	}
	if dataLen > 0 {
		if err := binary.Write(buffer, binary.BigEndian, data); err != nil {
			s := fmt.Sprintf("Pack data error , %v", err)
			return nil, errors.New(s)
		}
	}
	return buffer.Bytes(), nil
}

// ClientDecode :
func ClientDecode(rawConn net.Conn) (*protocol.NetDataProtocol, error) {
	newPackage := protocol.NetDataProtocol{}
	headData := make([]byte, protocol.DefaultHeadLength)
	n, err := io.ReadFull(rawConn, headData)
	if n != protocol.DefaultHeadLength {
		return nil, err
	}
	bytesBuffer := bytes.NewBuffer(headData)
	binary.Read(bytesBuffer, binary.BigEndian, &newPackage.CommandId)
	binary.Read(bytesBuffer, binary.BigEndian, &newPackage.DataLength)
	if newPackage.DataLength < 1 {
		return &newPackage, nil
	}
	data := make([]byte, newPackage.DataLength)
	dataNum, err2 := io.ReadFull(rawConn, data)
	if uint32(dataNum) != newPackage.DataLength {
		s := fmt.Sprintf("read data error, %v", err2)
		return nil, errors.New(s)
	}
	newPackage.Data = data
	return &newPackage, nil
}
