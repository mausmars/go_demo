package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"strconv"
	"sync"
	"time"
)

var Mutex sync.Mutex

var i = 0

func name2(s string) {
	Mutex.Lock()
	str := fmt.Sprint(windows.GetCurrentThreadId())
	fmt.Println("i==" + strconv.Itoa(i) + "  belong thread id " + str)
	i++
	defer Mutex.Unlock()

}

func main() {
	for i := 1; i <= 10; i++ {
		go name2(strconv.Itoa(i))
	}
	// 避免程序过快直接结束
	time.Sleep(100 * time.Second)
}
