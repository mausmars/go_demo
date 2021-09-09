package main

import (
	"fmt"
	"go_demo/src/sync_demo/module"
	"time"
)

func syncTest1() {
	times := 1000000
	totalTime := int64(0)

	var s *module.Hello
	st := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		s = &module.Hello{A: 1}
		module.Say(s)
	}
	totalTime += time.Now().UnixNano() - st

	time := totalTime / int64(times)
	fmt.Printf("new 调用次数=%d 平均时间= %d ns \n",times, time)
}

func syncTest2() {
	times := 1000000
	totalTime := int64(0)

	var s *module.Hello
	st := time.Now().UnixNano()
	for i := 0; i < times; i++ {
		s = module.Pool.Get().(*module.Hello)
		s.A = 1
		module.Say(s)
		module.Pool.Put(s)
	}
	totalTime += time.Now().UnixNano() - st

	time := totalTime / int64(times)
	fmt.Printf("pool 调用次数=%d 平均时间= %d ns \n",times, time)
}

func main() {
	syncTest2()
	syncTest1()
}
