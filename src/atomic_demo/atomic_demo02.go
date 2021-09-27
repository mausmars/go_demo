package main

import (
	"go.uber.org/atomic"
)

//原子加加
func main() {
	id := atomic.NewInt32(0)
	id.Inc()
}
