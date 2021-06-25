package main

import (
	"runtime"
	"time"
)

func main() {
	i := 0
	size := 1000000
	as := make([]*int, 0)
	for ; ; {
		a := getValue1()
		as = append(as, a)
		i++
		if i > size {
			break
		}
	}
	//a:=getValue1()
	//b:=getValue2()
	//println(*a)
	//println(b)
	runtime.GC()
	println("gc!")
	for i, a := range as {
		if *a != 1 {
			println(i, " ", a, " ", *a)
		}
	}
	println("over!")
	time.Sleep(10 * time.Second)
}

func getValue1() *int {
	i := 1
	return &i
}

func getValue2() int {
	i := 2
	return i
}
