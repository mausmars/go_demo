package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

/**
同一个方法里打印的线程id不一致。。。

逻辑cpu数量:12
主协程所属线程id =23236
协程_3 同一次执行 线程也不同。。。 count= 141855
协程_10 同一次执行 线程也不同。。。 count= 156698
协程_12 同一次执行 线程也不同。。。 count= 320174
协程_13 同一次执行 线程也不同。。。 count= 321676

 */
func GetCurrentThreadId() int {
	var user32 *syscall.DLL
	var GetCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		fmt.Printf("syscall.LoadDLL fail: %v\n", err.Error())
		return 0
	}
	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		fmt.Printf("user32.FindProc fail: %v\n", err.Error())
		return 0
	}

	var pid uintptr
	pid, _, err = GetCurrentThreadId.Call()

	return int(pid)
}

func name(s string, ch chan string) {
	//threadId := ""
	count := 0
	for {
		//time.Sleep(time.Second)
		t1 := windows.GetCurrentThreadId()
		t2 := windows.GetCurrentThreadId()
		if t1 != t2 {
			fmt.Println("协程_"+s+" 同一次执行 线程也不同。。。 count=", count)
		}

		//t1 := fmt.Sprint(GetCurrentThreadId())
		//if threadId == "" {
		//	threadId = t1
		//} else {
		//	if threadId != t1 {
		//		//fmt.Println("协程_" + s + " 线程id变更 " + threadId + "->" + t1)
		//		threadId = t1
		//	}
		//}

		//为了演示起来方便 我们每个协程都是相隔一秒才打印，否则命令行中刷起来太快，不好看执行过程
		//ch <- threadId
		//fmt.Println("协程_" + s + " belong threadId " + str)
		count++
	}
}

func record(ch chan string) {
	threadIds := map[string]string{}
	threadId := ""
	for {
		select {
		case threadId = <-ch:
			threadIds[threadId] = threadId
			//fmt.Println("线程数 " + strconv.Itoa(len(threadIds)))
		}
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	//逻辑cpu数量为4，代表我这个go程序 有4个p可以使用。每个p都会被分配一个系统线程。
	//这里因为我电脑的cpu是i5 4核心的，所以这里返回的是4. 如果你的机器是i7 四核心的，那这里返回值就是8了
	//因为intel的i7 cpu 有超线程技术，简单来说就是一个cpu核心 可以同时运行2个线程。
	fmt.Println("逻辑cpu数量:" + strconv.Itoa(runtime.NumCPU()))
	str := fmt.Sprint(windows.GetCurrentThreadId())
	fmt.Println("主协程所属线程id =" + str)

	ch := make(chan string)
	go record(ch)

	//既然在我机器上golang默认是4个逻辑线程，那我就将同步任务扩大到10个，看看执行结果
	for i := 1; i <= 14; i++ {
		go name(strconv.Itoa(i), ch)
	}
	// 避免程序过快直接结束
	time.Sleep(100 * time.Second)
}
