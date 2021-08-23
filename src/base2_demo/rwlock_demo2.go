package main

import (
	"runtime"
	"sync"
	"time"
)

var m2 *sync.RWMutex

func main() {
	//多个线程和单个线程效果不一样。多个可以并行read
	runtime.GOMAXPROCS(5)
	//runtime.GOMAXPROCS(1)

	m2 = new(sync.RWMutex)
	//写的时候啥都不能干
	go write2(1)
	go read2(2)
	go write2(3)
	go read2(4)
	time.Sleep(4 * time.Second)
}

func read2(i int) {
	m2.RLock()
	println(i, "read start")
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read end")
	m2.RUnlock()
}

func write2(i int) {
	m2.Lock()
	println(i, "write start")
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write end")
	m2.Unlock()

}
