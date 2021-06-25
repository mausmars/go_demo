package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type ILockService interface {
	Lock(key string)
	TryLock(key string) (ok bool, err error)
	Unlock(key string) (err error)
	//AddTimeout(key string, exTime int64) (ok bool, err error)
}

//---------------------------------------------------------------
const Lock_Value = "lock"

type RedisLockService struct {
	pool        *redis.Pool
	lockTimeout float32 //秒
	sleepTime   int32   //毫秒
}

func (s *RedisLockService) TryLock(key string) (ok bool, err error) {
	conn := s.pool.Get()
	defer conn.Close()

	key = s.createKey(key)
	str, err := redis.String(conn.Do("SET", key, Lock_Value, "PX", s.lockTimeout, "NX"))
	fmt.Println(str)
	if err == redis.ErrNil {
		// The lock was not successful, it already exists.
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *RedisLockService) Lock(key string) {
	for ; ; {
		ok, _ := s.TryLock(key)
		if ok {
			break
		} else {
			//runtime.Gosched()
			time.Sleep(time.Duration(s.sleepTime) * time.Millisecond)
		}
	}
}

func (s *RedisLockService) Unlock(key string) (err error) {
	conn := s.pool.Get()
	defer conn.Close()

	key = s.createKey(key)
	_, err = conn.Do("del", key)
	return
}

func (s *RedisLockService) createKey(key string) string {
	return fmt.Sprintf("redislock:%s", key)
}

//---------------------------------------------------------------
func CreateRedisLockService(address string) ILockService {
	if address == "" {
		address = "localhost:6379"
	}
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1200, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	s := &RedisLockService{
		pool:        pool,
		lockTimeout: 1500,
		sleepTime:   100,
	}
	return s
}

func main() {
	LockService := CreateRedisLockService("")
	key1 := "lock_1"
	key2 := "lock_2"
	key3 := "lock_3"

	go func() {
		fmt.Println("g1 准备获取锁")
		LockService.Lock(key1)
		LockService.Lock(key2)
		fmt.Println("g1 获取得锁  开始等待")
		time.Sleep(5 * time.Second)
		defer func() {
			LockService.Unlock(key1)
			LockService.Unlock(key2)
		}()
		fmt.Println("g1 釋放鎖 ")
	}()
	//time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("g2 准备获取锁")
		LockService.Lock(key1)
		LockService.Lock(key3)
		fmt.Println("g2 获取得锁  开始等待")
		time.Sleep(5 * time.Second)
		fmt.Println("g2 获取得锁 ")
		defer func() {
			LockService.Unlock(key1)
			LockService.Unlock(key3)
		}()
		fmt.Println("g2 釋放鎖 ")
	}()
	fmt.Println("主線程等待")
	time.Sleep(999999999 * time.Second)
}
