package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		//runtime.Gosched()用于让出CPU时间片
		runtime.Gosched()
		fmt.Println(s)
	}
}
//runtime.GOMAXPROCS(n)指定使用多核，不指定都是在一個線程下
//当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞

func main() {
	runtime.GOMAXPROCS(5)
	go say("world")
	say("hello")
}
