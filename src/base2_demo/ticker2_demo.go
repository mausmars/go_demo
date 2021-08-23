package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		<-t.C
		fmt.Println("timeout...")
		// need reset
		t.Reset(time.Second * 2)
	}
}
