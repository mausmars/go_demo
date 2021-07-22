package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
#include <stdio.h>
#include <unistd.h>
#include <sys/syscall.h>

int getTid() {
	return syscall(SYS_gettid);
}
*/
import "C"

func name(cid string) {
	count := 0
	for {
		t1 := C.getTid()
		t2 := C.getTid()
		if t1 != t2 {
			fmt.Println("协程_"+cid+" 同一次执行 线程也不同。。。 count=", count)
		}
		count++
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	//逻辑cpu数量为4，代表我这个go程序 有4个p可以使用。每个p都会被分配一个系统线程。
	//这里因为我电脑的cpu是i5 4核心的，所以这里返回的是4. 如果你的机器是i7 四核心的，那这里返回值就是8了
	//因为intel的i7 cpu 有超线程技术，简单来说就是一个cpu核心 可以同时运行2个线程。
	fmt.Println("逻辑cpu数量:" + strconv.Itoa(runtime.NumCPU()))
	str := fmt.Sprint(C.getTid())
	fmt.Println("主协程所属线程id =" + str)

	//既然在我机器上golang默认是4个逻辑线程，那我就将同步任务扩大到10个，看看执行结果
	for i := 1; i <= 14; i++ {
		go name(strconv.Itoa(i))
	}
	// 避免程序过快直接结束
	time.Sleep(100 * time.Second)
}
