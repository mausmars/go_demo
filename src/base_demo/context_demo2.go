package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//每个协程 context文内容继承父一级的，并能修改自己的，不影响其他协程上下文内容。

func main() {
	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, 1, "this is value a")
	ctx = context.WithValue(ctx, 2, "this is value b")
	fmt.Printf("value: %v\n", ctx.Value(1))

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(id int, ctx context.Context) {
			if id == 1 {
				ctx = context.WithValue(ctx, 2, "this is value b1")
			} else if id == 2 {
				ctx = context.WithValue(ctx, 2, "this is value b2")
			}

			for k := 0; k < 8; k++ {
				if id == 3 {
					ctx = context.WithValue(ctx, 2, "this is value b3")
				}
				fmt.Printf("id=%d, %s, %s \n", id, ctx.Value(1), ctx.Value(2))
				time.Sleep(1 * time.Second)
			}
			wg.Done()
		}(i, ctx)
	}
	time.Sleep(3 * time.Second)
	ctx = context.WithValue(ctx, 2, "this is value c")
	wg.Wait()
	fmt.Printf("main %s, %s \n", ctx.Value(1), ctx.Value(2))
}
