package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//go test -v -bench=. -benchmem -run=none

func BenchmarkAtomicDemo(b *testing.B) {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("atomic 花费时间：%d ns, count的值为：%d \n", time.Now().Sub(t).Nanoseconds(), count)
}

func BenchmarkMutexDemo(b *testing.B) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	count := int64(0)
	t := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			count++
			wg.Done()
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("mutex 花费时间：%d ns, count的值为：%d \n", time.Now().Sub(t).Nanoseconds(), count)
}
