package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer has expired.")
	}()
//timer.Stop() 这个接口从设计的时候就设计成了并不去关闭Channel
	timer.Stop()
	//c:=(chan time.Time)(timer.C)
	//close(timer.C)
	//timer.Reset(0  * time.Second)
	fmt.Println("Timer Stop")
	time.Sleep(10 * time.Second)
	fmt.Println("over")
}
