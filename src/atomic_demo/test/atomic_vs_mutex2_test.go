package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//go test -v -cpu 1,2,4 -benchmem -bench=. atomic_vs_mutex2_test.go

func BenchmarkAtomic(b *testing.B) {
	var number int32
	fmt.Println("n=", b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			atomic.AddInt32(&number, 1)
		}()
	}
}

func BenchmarkMutex(b *testing.B) {
	var number int
	lock := sync.Mutex{}

	fmt.Println("n=", b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer lock.Unlock()
			lock.Lock()
			number++
		}()
	}
}
