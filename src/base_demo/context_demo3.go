package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(999 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("error ", ctx.Err()) // prints "context deadline exceeded"
	}
}
