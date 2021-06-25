package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(5)
	for i:=0;i<2;i++{
		go func() {
			fmt.Println("test")
			time.Sleep(10*time.Second)
		}()
	}

	fmt.Println("version ",runtime.Version())
	fmt.Println("cpu ",runtime.NumCPU())
	fmt.Println("CgoCall ",runtime.NumCgoCall())
	fmt.Println("Goroutine ",runtime.NumGoroutine())
}
