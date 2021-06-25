package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	type key = struct{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, 1, "this is value")
	fmt.Printf("value: %v\n", ctx.Value(1))

	for i:=0;i<10;i++{
		go func(ctx context.Context){
			ctx = context.WithValue(ctx, 2, "this is value")


			fmt.Println(ctx.Value(1))
		}(ctx)
	}


	time.Sleep(10000*time.Millisecond)
}


