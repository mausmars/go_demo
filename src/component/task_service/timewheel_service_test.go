package task_service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomTime(num int) time.Duration {
	return time.Duration(rand.Intn(num)) * time.Second
}

func TestMatchService(t *testing.T) {
	callback := func(userData interface{}) {
		// do something
		fmt.Println("task callback  userData=", userData)
	}

	tw := NewTimeWheelService(1*time.Second, 3600, callback)
	// 启动时间轮
	tw.Startup()
	//test1(tw)
	test2(tw)
}

func test2(tw *TimeWheelService) {
	tw.InsertTask(1*time.Second, 1, 1)
	time.Sleep(20 * time.Millisecond)
	tw.InsertTask(2*time.Second, 1, 2)
	time.Sleep(20 * time.Millisecond)
	tw.RemoveTask(1)
	tw.RemoveTask(1)
	time.Sleep(20 * time.Millisecond)

	time.Sleep(40 * time.Second)
}

func test1(tw *TimeWheelService) {
	for i := 0; i < 20; i++ {
		time := randomTime(25)
		tw.InsertTask(time, i, i)
	}
	time.Sleep(30 * time.Second)
	// 删除定时器, 参数为添加定时器传递的唯一标识
	// 停止时间轮
	tw.Shutdown()
	//select{}
	time.Sleep(40 * time.Second)
}
