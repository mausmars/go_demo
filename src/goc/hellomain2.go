package main

import "C"
import (
	"fmt"
	"time"
)

func main() {
	//s := "Hello"
	currentTime := time.Now()

	for i := 0; i < 10000; i++ {
		//fmt.Println(s)
		c := 1 + 2
		fmt.Println(c)
	}
	time := time.Now().Sub(currentTime).Nanoseconds()
	fmt.Println("time=", time/1000, " ws")
}
