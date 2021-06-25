package main

import (
	"go_demo/src/component/statistics_service"
	protomsg "go_demo/src/gin_demo/httpserver/proto"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
)

func main() {
	statisticsService := &statistics_service.StatisticsService{
	}
	statisticsService.StartUp()

	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}

	for i := 0; i < 10000; i++ {
		go func(pool *redis.Pool, id int) {
			client := pool.Get()
			key := "key_" + strconv.Itoa(id)
			{
				msg := &protomsg.Command{
					CommandId: int32(id),
					Body:      []byte("abcdefghijklmnopqrstuvwxyz"),
				}
				data, _ := proto.Marshal(msg)

				commandInfo := &statistics_service.CommandInfo{
					SendTime: time.Now(),
					Command:  "Redis Save",
				}
				_, err := client.Do("PSETEX", key, 100000, data)
				if err != nil {
					//println(err.Error())
					commandInfo.IsSuccess = false
					statisticsService.Recorder(commandInfo)
				} else {
					//println(reply.(string))
					commandInfo.IsSuccess = true
					statisticsService.Recorder(commandInfo)
				}
			}
			//time.Sleep(2*time.Second)
			{
				commandInfo := &statistics_service.CommandInfo{
					SendTime: time.Now(),
					Command:  "Redis Get",
				}
				reply, err := client.Do("GET", key)
				if err != nil {
					//println(err.Error())
					commandInfo.IsSuccess = false
					statisticsService.Recorder(commandInfo)
				} else {
					if reply == nil {
						//println(reply)
						commandInfo.IsSuccess = true
						statisticsService.Recorder(commandInfo)
					} else {
						commandInfo.IsSuccess = true
						statisticsService.Recorder(commandInfo)

						msg := &protomsg.Command{}
						proto.Unmarshal(reply.([]uint8), msg)
						//println(msg.CommandId)
						//println(string(msg.Body))
					}
				}
			}
			defer client.Close()
		}(pool, i)
	}

	var i int
	var explain = "----------------------------\n" +
		"输入指令序号\n" +
		"1.退出测试\n" +
		"2.查看测试信息"
	for {
		fmt.Println(explain)
		fmt.Scan(&i)
		var isOver = false
		switch i {
		case 1:
			statisticsService.Shutdown()
			isOver = true
			break
		case 2:
			statisticsService.Show()
			break
		}
		if isOver {
			break
		}
	}
}
