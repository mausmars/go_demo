package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg  sync.WaitGroup
	mu  sync.Mutex
	num int32
)

func wrap(callback func()) {
	wg.Add(1)
	go func() {
		callback()
		wg.Done()
	}()
}

// go build -race
func main() {
	// 1099 、892, 并发不安全
	//wrap(incNumAtomic)
	//wrap(incNumAtomic)

	// 1200, 并发安全
	wrap(incNumAtomic2)
	wrap(incNumAtomic2)

	wg.Wait()
	fmt.Printf("num=%d\n", num)
}

func incNumAtomic() {
	for i := 0; i < 600; i++ {
		// atomic.Load*系列函数只能保证读取的不是正在写入的值（比如只被修改了一半的数据）
		// 同样的atomic.Store* 只保证写入是原子操作(保证写入操作的完整性)
		val := atomic.LoadInt32(&num)
		atomic.StoreInt32(&num, val+1)
	}
}

func incNumAtomic2() {
	for i := 0; i < 600; i++ {
		atomic.AddInt32(&num, 1)
	}
}
