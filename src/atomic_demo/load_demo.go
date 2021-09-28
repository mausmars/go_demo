package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type L struct{ x, y, z int }
	var l1 = L{9, 10, 11}
	var V atomic.Value
	V.Store(l1)

	var l2 = V.Load()
	if l2 != nil {
		l2 = l2.(L)
		fmt.Println(l2)
		fmt.Println(l1 == l2)
	} else {
		fmt.Println("load is nil")
	}
}
