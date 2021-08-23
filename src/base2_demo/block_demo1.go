package main

import (
	"fmt"
	"time"
)

func main() {
	size := 2

	chs := make([]chan int, size)
	for i := 0; i < size; i++ {
		chs[i] = make(chan int)
		go func(a int, i int, ch chan int) {
			fmt.Println("sleep start!", a, i)
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Println("sleep over!", a, i)
			ch <- 1
		}(1, i, chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("over!")
}
