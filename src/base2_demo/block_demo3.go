package main

import (
	"fmt"
	"runtime"
	"time"
)

//让当前协程永久阻塞

func DoSomething(i int) {
	for {
		// 做点什么...
		fmt.Println("sleep start!", i)
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("sleep over!", i)
		runtime.Gosched() // 防止本协程霸占CPU不放
	}
}

func main() {
	go DoSomething(1)
	go DoSomething(2)
	select {}

	fmt.Println("over!")
}
