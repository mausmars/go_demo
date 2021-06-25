package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig 为SIGINT中断设置一个监听器
func CatchSig(ch chan os.Signal, done chan bool) {
	// 在等待信号时阻塞
	sig := <-ch
	// 当接收到信号时打印
	fmt.Println("\nsig received:", sig)

	// 对信号类型进行处理
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}
	// 终止
	done <- true
}

func main() {
	// 初始化通道
	signals := make(chan os.Signal)
	done := make(chan bool)

	// 将它们连接到信号lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 如果一个信号被这个go例程捕获，它将写入 done
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	// 程序会持续打印日志直到done通道被写入
	<-done
	fmt.Println("Done!")

}