package main

import (
	"fmt"
	"github.com/modern-go/gls"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		go func() {
			gls.ResetGls(gls.GoID(), make(map[interface{}]interface{},10))
			gls.Set("goid", gls.GoID())
			fmt.Println(gls.Get("goid"))
		}()
	}
	time.Sleep(1000 * time.Millisecond)
}
