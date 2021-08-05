package main

/*
#include "cgo_demo.h"
*/
import "C"

import "fmt"

func GoSum(a, b int) {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)
}

func main() {
	GoSum(4, 5)
}
