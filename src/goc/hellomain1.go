package main

/*
#include "hello.c"
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	s := "Hello"
	cs := C.CString(s)

	currentTime := time.Now()
	for i := 0; i < 10000; i++ {
		c := C.sum(C.int(1), C.int(2), cs)
		fmt.Println(c)
	}
	time := time.Now().Sub(currentTime).Nanoseconds()
	fmt.Println("time=", time/1000, " ws")
	defer C.free(unsafe.Pointer(cs))
}
