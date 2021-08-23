package main

import (
	"fmt"
	"runtime"
)

func say2(s string) {
	for i := 0; i < 2; i++ {
		//runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(10)
	go say2("world")
	//runtime.Gosched()
	say2("hello")
}
