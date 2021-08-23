package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	size := 2
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(a int, i int, done func()) {
			defer done()
			fmt.Println("sleep start!", a, i)
			time.Sleep(time.Duration(2) * time.Second)
			fmt.Println("sleep over!", a, i)
		}(i, 1, wg.Done)
	}
	wg.Wait()
	fmt.Println("over!")
}
