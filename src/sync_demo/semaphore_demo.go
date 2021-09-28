package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"math/rand"
	"runtime"
	"time"
)

//总结
//Acquire和TryAcquire都可用于获取资源，Acquire是可以阻塞的获取资源，TryAcquire只能非阻塞的获取资源;
//Release对于waiters的唤醒原则，总是先进先出，避免资源需求比较大的waiter被饿死;

// Example_workerPool演示如何使用信号量来限制
// 用于并行任务的goroutine。
//---------------------------
//信号量是在并发编程中比较常见的一种同步机制，它会保证持有的计数器在0到初始化的权重之间，每次获取资源时都会将信号量中的计数器减去对应的数值，
//在释放时重新加回来，当遇到计数器大于信号量大小时就会进入休眠等待其他进程释放信号。

//go中的semaphore，提供sleep和wakeup原语，使其能够在其它同步原语中的竞争情况下使用。当一个goroutine需要休眠时，将其进行集中存放，
//当需要wakeup时，再将其取出，重新放入调度器中。

//go中本身提供了semaphore的相关方法，不过只能在内部调用

func main() {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	var (
		maxWorkers = runtime.GOMAXPROCS(0)                    //12
		sem        = semaphore.NewWeighted(int64(maxWorkers)) //权重12
		out        = make([]int, 32)
	)
	// Compute the output using up to maxWorkers goroutines at a time.
	for i := range out {
		// When maxWorkers goroutines are in flight, Acquire blocks until one of the workers finishes.
		//如果 值大于 maxWorkers，Acquire 阻塞，等待释放
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		go func(i int) {
			defer sem.Release(1) //释放
			// doSomething
			out[i] = i + 1
			fmt.Println("id =", i, " doSomething!")
			n := rand.Intn(5) + 3
			time.Sleep(time.Duration(n) * time.Second)
		}(i)
	}
	// Acquire all of the tokens to wait for any remaining workers to finish.
	//
	// If you are already waiting for the workers by some other means (such as an
	// errgroup.Group), you can omit this final Acquire call.
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}
	fmt.Println(out)
}
