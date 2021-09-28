package main

import (
	"fmt"
	"sync"
)

func main() {
	oc := &sync.Once{}

	for i := 0; i < 10; i++ {
		oc.Do(func() {
			fmt.Println("only once")
		})
	}
}
