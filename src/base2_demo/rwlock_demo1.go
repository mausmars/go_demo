package main

import (
	"runtime"
	"sync"
	"time"
)

var m1 *sync.RWMutex

func main() {
	runtime.GOMAXPROCS(5)

	m1 = new(sync.RWMutex)
	//可以多个同时读
	go read1(1)
	go read1(2)
	time.Sleep(2 * time.Second)
}

func read1(i int) {
	println(i, "read start")
	m1.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m1.RUnlock()
	println(i, "read end")
}
