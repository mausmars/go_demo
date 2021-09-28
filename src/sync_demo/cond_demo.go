package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var status int64
	wg := sync.WaitGroup{}
	n := 10
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < n; i++ {
		wg.Add(1)
		go listen(i, c, &wg, &status)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c, n, &status)

	wg.Wait()
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, os.Interrupt)
	//<-ch
}

func broadcast(c *sync.Cond, n int, status *int64) {
	c.L.Lock()
	atomic.StoreInt64(status, 1)
	//c.Broadcast() // 唤醒
	for i := 0; i < n; i++ {
		c.Signal() //唤醒一个，一共唤醒n次
	}
	c.L.Unlock()
}

func listen(id int, c *sync.Cond, wg *sync.WaitGroup, status *int64) {
	c.L.Lock()
	for atomic.LoadInt64(status) != 1 {
		c.Wait() //wait
	}
	fmt.Println(id, " listen")
	c.L.Unlock()

	wg.Done()
}
