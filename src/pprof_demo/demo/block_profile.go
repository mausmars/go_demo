package main

import (
	"fmt"
	"go_demo/src/profile"
	"sync"
	"time"
)

func main() {
	defer profile.Start(profile.BlockProfile).Stop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(done func()) {
		defer done()
		fmt.Println("sleep start!")
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("sleep over!")
	}(wg.Done)
	wg.Wait()
	fmt.Println("over!")
}
