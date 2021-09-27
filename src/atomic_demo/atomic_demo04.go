package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func atomicDemo1() {
	var wg sync.WaitGroup

	// Declaring u, any pointer
	var u uintptr

	// For loop
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// Function with
		// AddUintptr method
		go func() {
			atomic.AddUintptr(&u, 1)
			wg.Done()
		}()
	}
	wg.Wait() //阻塞直到所有任务完成
	// Prints loaded values address
	fmt.Println(atomic.LoadUintptr(&u))
	fmt.Println("over")
}

func atomicDemo2() {
	var wg sync.WaitGroup

	// Declaring u
	var u int

	// For loop
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// Function with
		// AddUintptr method
		go func() {
			u++
			wg.Done()
		}()
	}
	wg.Wait() //阻塞直到所有任务完成
	// Prints loaded values address
	fmt.Println(u)
	fmt.Println("over")
}

func main() {
	runtime.GOMAXPROCS(10)
	atomicDemo1()
	//atomicDemo2()
}
