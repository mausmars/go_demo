package main

import (
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex

	go func() {
		lock.Lock()
		time.Sleep(1000000 * time.Second)
		lock.Unlock()
	}()

	go func() {
		lock.Lock()
		time.Sleep(1 * time.Second)
		lock.Unlock()
	}()

	time.Sleep(1000000 * time.Second)
}
