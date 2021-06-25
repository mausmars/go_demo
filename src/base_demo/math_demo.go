package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	startTime := time.Now()
	for i := 0; i < 1000000; i++ {
		math.Sqrt(9999)
	}
	fmt.Println("time ",(time.Now().UnixNano() - startTime.UnixNano())/1000000," ms")
}
