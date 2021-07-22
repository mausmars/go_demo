package main

import (
	"runtime"
	"time"
)

func main() {
	println("NumCPU ", runtime.NumCPU())

	for i := 0; i <= 20; i++ {
		go func() {
			time.Sleep(time.Duration(2) * time.Second)
		}()
	}
	println("NumGoroutine ", runtime.NumGoroutine())
}
