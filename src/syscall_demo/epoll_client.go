package main

import (
	"fmt"
	"syscall"
)

func main() {
	server := &EpollClientService{
		Port: 8000,
		Addr: [4]byte{127, 0, 0, 1},
	}
	server.StartUp()
}

const (
	MaxDataSize_Client = 32 * 1024
)


type EpollClientService struct {
	Port   int
	Addr   [4]byte
	sockfd int
}

func (s *EpollClientService) StartUp() {
	var err error
	s.sockfd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	err = s.setSockeopt()
	if err != nil {
		fmt.Println("setSockeopt error: ", err.Error())
		return
	}
	serverAddr := &syscall.SockaddrInet4{
		Port: s.Port,
		Addr: s.Addr,
	}
	err = syscall.Connect(s.sockfd, serverAddr)
	if err != nil {
		fmt.Println("Connect error: ", err.Error())
		return
	}
	msg := "Hello, world!"
	_, err = syscall.Write(s.sockfd, []byte(msg))
	if err != nil {
		fmt.Println("Write error: ", err.Error())
		return
	}
	// 数据缓存区域
	buf := make([]byte, MaxDataSize_Client)
	_, err = syscall.Read(s.sockfd, buf)
	if err != nil {
		fmt.Println("Read error: ", err.Error())
		return
	}
	fmt.Println("received msg=", string(buf))
	//关闭连接
	syscall.Close(s.sockfd)
}

func (s *EpollClientService) setSockeopt() error {
	isReuseaddr := 1
	err := syscall.SetsockoptInt(s.sockfd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, isReuseaddr)
	if err != nil {
		//closesocket（一般不会立即关闭而经历TIME_WAIT的过程）后想继续重用该socket
		fmt.Errorf("setsockopt SO_REUSEADDR error! \n")
		return err
	}
	/**
	  选项              间隔      关闭方式  等待关闭与否
	  SO_DONTLINGER   不关心     优雅         否
	  SO_LINGER       零        强制         否
	  SO_LINGER       非零      优雅         是
	*/
	linger := &syscall.Linger{
		Onoff:  1, //(在closesocket()调用,但是还有数据没发送完毕的时候容许逗留 如果m_sLinger.l_onoff=0;则功能和2作用相同
		Linger: 5, //(容许逗留的时间为5秒)
	}
	err = syscall.SetsockoptLinger(s.sockfd, syscall.SOL_SOCKET, syscall.SO_LINGER, linger)
	if err != nil {
		//大致意思就是说SO_LINGER选项用来设置当调用closesocket时是否马上关闭socket
		fmt.Errorf("setsockopt SO_LINGER error! \n")
		return err
	}
	//发送数据的时，希望不经历由系统缓冲区到socket缓冲区的拷贝而影响程序的性能
	//    int zero=0;
	//    setsockopt(sockfd,SOL_SOCKET,SO_RCVBUF,(const void *)&zero,sizeof(zero));
	//    setsockopt(sockfd,SOL_SOCKET,SO_SNDBUF,(const void *)&zero,sizeof(zero));
	recvbuf := 32 * 1024 //设置为32K
	err = syscall.SetsockoptInt(s.sockfd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, recvbuf)
	if err != nil {
		// 接收缓冲区
		fmt.Errorf("setsockopt SO_RCVBUF error! \n")
		return err
	}
	sendbuf := 32 * 1024 //设置为32K
	err = syscall.SetsockoptInt(s.sockfd, syscall.SOL_SOCKET, syscall.SO_SNDBUF, sendbuf)
	if err != nil {
		//发送缓冲区
		fmt.Errorf("setsockopt SO_SNDBUF error! \n")
		return err
	}
	isTcpnodelay := 1
	err = syscall.SetsockoptInt(s.sockfd, syscall.SOL_SOCKET, syscall.TCP_NODELAY, isTcpnodelay)
	if err != nil {
		//发送缓冲区
		fmt.Errorf("setsockopt TCP_NODELAY error! \n")
		return err
	}
	return nil
}

func (s *EpollClientService) setNonBlocking() error {
	return syscall.SetNonblock(s.sockfd, true)
}
