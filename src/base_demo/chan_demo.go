package main

import (
	"fmt"
	"time"
)

type ChanDemo struct {
	commandChan chan int32
}

func (r *ChanDemo) Run() {
	go func() {
		for {
			select {
			case command := <-r.commandChan:
				fmt.Println("command=", command)
				time.Sleep(1 * time.Second)
				go func() {
					r.commandChan <- command + 1
				}()
				break
			default:
				break
			}
		}
	}()
}

func main() {
	chanDemo := &ChanDemo{
		//commandChan: make(chan int32,10),
		commandChan: make(chan int32),
	}
	chanDemo.Run()

	chanDemo.commandChan <- 1

	time.Sleep(99 * time.Second)
}
