package main

import (
	"fmt"
	"go.uber.org/atomic"
)

func main() {
	isRuning := atomic.NewBool(false)

	isSuccess := isRuning.CAS(false, true)
	fmt.Println(isSuccess)

	isSuccess = isRuning.CAS(false, true)
	fmt.Println(isSuccess)

	isSuccess = isRuning.CAS(false, true)
	fmt.Println(isSuccess)
}
