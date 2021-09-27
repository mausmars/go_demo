package main

import (
	"context"
	"fmt"
	"time"
)

//WithDeadline
func main() {
	d := time.Now().Add(5000 * time.Millisecond)
	// 设置上下文超时
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	//case <-time.After(1000 * time.Millisecond):
	case <-time.After(999 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		//上下文超时，取消
		fmt.Println("error ", ctx.Err())
	}
}
