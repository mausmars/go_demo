package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// Assigning values
	// to the uintptr
	var (
		i uintptr = 98
		j uintptr = 255
		k uintptr = 6576567667788
		l uintptr = 5
	)

	// Calling LoadUintptr method
	// with its parameters
	load1 := atomic.LoadUintptr(&i)
	load2 := atomic.LoadUintptr(&j)
	load3 := atomic.LoadUintptr(&k)
	load4 := atomic.LoadUintptr(&l)

	// Displays uintptr value
	// loaded in the *addr
	fmt.Println(load1)
	fmt.Println(load2)
	fmt.Println(load3)
	fmt.Println(load4)
}
