package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

func test1() {
	/*	   DialTimeout time.Duration `json:"dial-timeout"`	   Endpoints []string `json:"endpoints"`	*/
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
}

func test2() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	//操作完毕，取消etcd
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	//取值，设置超时为1秒
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func test3() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	go func() {
		n := 4
		defer cli.Close()
		for i := 0; i < 10; i++ {
			if i%n == 0 {
				cli.Put(context.Background(), "/logagent/conf/", "8888888")
			} else if i%n == 1 {
				resp, err := cli.Get(context.Background(), "/logagent/conf/")
				if err == nil {
					for _, ev := range resp.Kvs {
						fmt.Printf("%s : %s\n", ev.Key, ev.Value)
					}
				}else{
					fmt.Println("connect failed, err:", err)
				}
			} else if i%n == 2 {
				cli.Delete(context.Background(), "/logagent/conf/")
			} else {
				resp, err := cli.Get(context.Background(), "/logagent/conf/")
				if err == nil {
					for _, ev := range resp.Kvs {
						fmt.Printf("%s : %s\n", ev.Key, ev.Value)
					}
				}else{
					fmt.Println("connect failed, err:", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		rch := cli.Watch(context.Background(), "/logagent/conf/")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}
