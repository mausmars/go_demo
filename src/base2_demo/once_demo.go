package main

// 官方案例

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) // 多次调用
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}