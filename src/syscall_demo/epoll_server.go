package main

import (
	"fmt"
	"syscall"
)

func main() {
	server := &EpollServerService{
		Port: 8000,
	}
	server.StartUp()
}

const (
	BackLog = 10000
	// Since Linux 2.6.8, the size argument is ignored, but must be greater than zero.
	MaxEvents = 50000

	MaxDataSize_Server = 32 * 1024

	ErrEvents = syscall.EPOLLERR | syscall.EPOLLHUP | syscall.EPOLLRDHUP
	OutEvents = syscall.EPOLLOUT | (syscall.EPOLLET & 0xffffffff)
	InEvents  = syscall.EPOLLIN | (syscall.EPOLLET & 0xffffffff)
)

type EpollServerService struct {
	Port int

	sockfd int
}

func (s *EpollServerService) StartUp() {
	var err error
	// 创建并监听tcp socket
	err = s.socketBind()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	// 设置socket为非阻塞
	err = s.setNonBlocking()
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	// 创建epoll句柄
	var epfd int
	epfd, err = syscall.EpollCreate1(0)
	if epfd == -1 {
		fmt.Println("epoll_create error ", err.Error())
		return
	}
	// epoll_ctl
	event := &syscall.EpollEvent{
		Fd:     int32(s.sockfd),
		Events: InEvents,
	}
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, s.sockfd, event)
	if err != nil {
		fmt.Println("epoll_ctl error ", err.Error())
		return
	}
	// Buffer where events are returned
	//var events [MaxEvents]syscall.EpollEvent
	//events := make([]syscall.EpollEvent, MaxEvents)

	// listen
	err = syscall.Listen(s.sockfd, BackLog)
	if err != nil {
		fmt.Println("listen error ", err.Error())
		return
	}
	var n int
	var connFd int

	for {
		var numbytes int
		events := make([]syscall.EpollEvent, 100)

		// 数据缓存区域
		buf := make([]byte, MaxDataSize_Server)

		n, err = syscall.EpollWait(epfd, events, -1)
		if err != nil {
			fmt.Println("epollwait error ", err.Error())
			break
		} else {
			fmt.Printf("epoll_wait 触发! n=%d \n", n)
		}
		for i := 0; i < n; i++ {
			/* We have a notification on the listening socket, which means one or more incoming connections. */
			if int(events[i].Fd) == s.sockfd {
				connFd, _, err = syscall.Accept(s.sockfd)
				if err != nil {
					fmt.Println("accept error ", err.Error())
					continue
				}
				fmt.Println("accept success!")

				syscall.SetNonblock(s.sockfd, true)
				//EPOLLIN 连接到达,有数据来临; EPOLLET 边缘触发
				event.Events = InEvents
				event.Fd = int32(connFd)
				//将新的fd添加到epoll的监听队列中
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, event)
				if err != nil {
					fmt.Println("epollctl error ", connFd, err)
					continue
				}
			} else if (events[i].Events & syscall.EPOLLIN) != 0 {
				//接收到数据，读socket
				connFd = int(events[i].Fd)
				if (connFd) < 0 {
					continue
				}
				numbytes, err = syscall.Read(connFd, buf)
				if err != nil {
					fmt.Println("Read error ", connFd, err)
					syscall.Close(connFd)
					event.Fd = -1
					syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, connFd, event)
					continue
				}
				if numbytes <= 0 {
					syscall.Close(connFd)
					event.Fd = -1
					syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, connFd, event)
					continue
				}
				fmt.Println("received msg=", string(buf))

				event.Events = OutEvents
				event.Fd = int32(connFd)
				//修改标识符，等待下一个循环时发送数据，异步处理的精髓
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_MOD, connFd, event)
				if err != nil {
					fmt.Println("epollctl error ", connFd, err)
					continue
				}
			} else if (events[i].Events & syscall.EPOLLOUT) != 0 {
				//有数据待发送，写socket
				connFd = int(events[i].Fd)
				syscall.Write(connFd, []byte("server resp")) //发送数据

				event.Events = InEvents
				event.Fd = int32(connFd)
				//修改标识符，等待下一个循环时接收数据
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_MOD, connFd, event)
			} else if (events[i].Events & ErrEvents) != 0 {
				// An error has occured on this fd, or the socket is not ready for reading (why were we notified then?)
				fmt.Println("epoll error ")
				connFd = int(events[i].Fd)
				syscall.Close(connFd)
				event.Fd = -1
				syscall.EpollCtl(epfd, syscall.EPOLL_CTL_DEL, connFd, event)
			}
		}
	}
}

func (s *EpollServerService) socketBind() error {
	var err error
	s.sockfd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return err
	}
	err = s.setSockeopt()
	if err != nil {
		fmt.Println("setSockeopt error: " + err.Error())
		return err
	}

	serverAddr := &syscall.SockaddrInet4{
		Port: s.Port,
		Addr: [4]byte{0, 0, 0, 0},
	}
	err = syscall.Bind(s.sockfd, serverAddr)
	if err != nil {
		syscall.Close(s.sockfd)
		fmt.Println("bind error: " + err.Error())
		return err
	}
	return nil
}

func (s *EpollServerService) setSockeopt() error {
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

func (s *EpollServerService) setNonBlocking() error {
	return syscall.SetNonblock(s.sockfd, true)
}
