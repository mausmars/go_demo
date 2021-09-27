package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)

	pid := sync.ProcPin()
	fmt.Printf("PID: %v \n", pid)
	sync.ProcUnpin()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			spid := sync.ProcPin()
			fmt.Printf("SPID: %v \n", spid)
			sync.ProcUnpin()
			wg.Done()
		}()
	}
	wg.Wait() //阻塞直到所有任务完成
	fmt.Println("over")
}
