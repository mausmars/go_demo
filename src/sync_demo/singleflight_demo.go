package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

//singleflight 减少重复操作，函数没执行完毕不会重复请求。

//singleflight 是 Go 语言扩展包中提供了另一种同步原语，它能够在一个服务中抑制对下游的多次重复请求。
//一个比较常见的使用场景是：我们在使用 Redis 对数据库中的数据进行缓存， 发生缓存击穿时，大量的流量都会打到数据库上进而影响服务的尾延时

// Do 执行函数, 对同一个 key 多次调用的时候，在第一次调用没有执行完的时候只会执行一次 fn 其他的调用会阻塞住等待这次调用返回 v, err 是传入的 fn 的返回值
//shared 表示是否真正执行了 fn 返回的结果，还是返回的共享的结果
//func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)

// DoChan 和 Do 类似，只是 DoChan 返回一个 channel，也就是同步与异步的区别
//func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result

// Forget 用于通知 Group 删除某个 key 这样后面继续这个 key 的调用的时候就不会在阻塞等待了
//func (g *Group) Forget(key string)

func getArticle(key string, count *int32, i int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	n := i % 2
	//atomic.AddInt32(count, 1)
	time.Sleep(time.Duration(n) * time.Millisecond)
	fmt.Println("Article ", key)
	return key, nil
}

func singleflightGetArticle(sg *singleflight.Group, key string, count *int32, i int) (string, error) {
	v, err, _ := sg.Do(key, func() (interface{}, error) {
		return getArticle(key, count, i)
	})
	return v.(string), err
}

func main() {
	var wg sync.WaitGroup

	n := 1000
	sg := &singleflight.Group{}
	var count int32

	key := fmt.Sprintf("%d", 1)
	now := time.Now()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetArticle(sg, key, &count, i)
			//res, _ := getArticle(key)
			if res != key {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}
