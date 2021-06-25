package main

import (
	"fmt"
	lock "github.com/bsm/redis-lock"
	"github.com/go-redis/redis"
	"time"
)

//https://godoc.org/github.com/bsm/redis-lock

func createLockService() *LockService {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	opts := &lock.Options{
		RetryCount:  0,
		LockTimeout: 1 * time.Second,
		RetryDelay:  500 * time.Millisecond,
	}
	s := &LockService{
		client: client,
		opts:   opts,
	}
	return s
}

type LockService struct {
	opts   *lock.Options
	client *redis.Client
}

func (s *LockService) createLock(lockKey string) *lock.Locker {
	locker := lock.New(s.client, lockKey, s.opts)
	return locker
}

func (s *LockService) Lock(lock *lock.Locker) {
	for ; ; {
		ok, _ := lock.Lock()
		if ok {
			break
		} else {
			//runtime.Gosched()
			time.Sleep(time.Duration(30) * time.Millisecond)
		}
	}
}
func (s *LockService) TryLock(lock *lock.Locker) (bool, error) {
	return lock.Lock()
}
func (s *LockService) UnLock(lock *lock.Locker) {
	lock.Unlock()
}

func main() {
	lockService := createLockService()

	go func() {
		fmt.Println("g1 准备获取锁")
		locker := lockService.createLock("test_1")
		lockService.Lock(locker)
		fmt.Println("g1 获取得锁  开始等待")
		time.Sleep(10 * time.Second)
		lockService.UnLock(locker)
		fmt.Println("g1 釋放鎖 ")
	}()
	time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("g2 准备获取锁")
		locker := lockService.createLock("test_1")
		lockService.Lock(locker)
		fmt.Println("g2 获取得锁 ")
		lockService.UnLock(locker)
		fmt.Println("g2 釋放鎖 ")
	}()
	fmt.Println("主線程等待")
	time.Sleep(999999999 * time.Second)
}
