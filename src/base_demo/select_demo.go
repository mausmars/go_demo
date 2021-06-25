package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c3 := make(chan int)
	c2 := make(chan int)

	go func (c1, c2, c3 chan int){
		i1 :=0
		i2 :=0

		for{
			select {
			case i1 = <-c1:
				fmt.Println("received ", i1, " from c1")
				i2++
			case c2 <- i2:
				fmt.Println("sent ", i2, " to c2")
			case i3, ok := (<-c3):  // same as: i3, ok := <-c3
				if ok {
					fmt.Println("received ", i3, " from c3")
				} else {
					fmt.Println("c3 is closed")
				}
			//default:
			//	fmt.Printf("no communication\n")
			}
		}
	}(c1,c2,c3)

	go func (c1, c2, c3 chan int){
		for i:=0;i<10;i++{
			fmt.Println("c2 1 read ",<-c2)
			//c1<-999
			//c3<-888
			//c1<-777
		}
	}(c1,c2,c3)

	go func (c1, c2, c3 chan int){
		for i:=0;i<10;i++{
			//c1<-777
			//c1<-999
			fmt.Println("c2 2 read ",<-c2)
			//c3<-888
		}
	}(c1,c2,c3)



	//c2<-777
	//close(c3)

	time.Sleep(90000*time.Millisecond)
}