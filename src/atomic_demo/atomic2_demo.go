package main

import (
	"go.uber.org/atomic"
)

func main() {
	id := atomic.NewInt32(0)

	id.Inc()
}
