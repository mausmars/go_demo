package inserttable_service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestInsertTableService(t *testing.T) {
	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	roomStateMgr := NewRoomStateMgr(pool)

	roomState := &RoomState{
		RoomServiceId: "localhost",
		Rid:           1,
		RoomId:        101,
		SeatIds:       []int32{1, 2}, //空位
	}
	// 预先加入一些空位
	roomStateMgr.saveRoomState(roomState)

	//启动插桌服务
	rids := []int32{1, 2}
	service := NewInsertTableService(rids, roomStateMgr)
	service.Startup()

	//获取可用的空位
	callback := func(isSuccess bool, seatId int32, roomState *RoomState) {
		if roomState != nil {
			fmt.Println("callback, isSuccess=", isSuccess, " seatId=", seatId, " roomId=", roomState.RoomId)
		} else {
			fmt.Println("callback, isSuccess=", isSuccess, " seatId=", seatId, " roomId=", 0)
		}
	}
	getSeatIdCommand := &GetSeatIdCommand{
		fpid:     10001,
		rid:      1,
		callback: callback,
	}
	service.GetSeatIdCommand(getSeatIdCommand)
	time.Sleep(5 * time.Second)
	service.GetSeatIdCommand(getSeatIdCommand)
	time.Sleep(5 * time.Second)
	service.GetSeatIdCommand(getSeatIdCommand)
	time.Sleep(5 * time.Second)

	//动态加入一些空位
	insertTableRoomCommand := &AddSeatIdCommand{
		roomServiceId: "localhost",
		rid:           1,
		roomId:        101,
		seatId:        2,
	}
	service.AddSeatIdCommand(insertTableRoomCommand)

	service.GetSeatIdCommand(getSeatIdCommand)
	time.Sleep(5 * time.Second)
	service.GetSeatIdCommand(getSeatIdCommand)

	time.Sleep(60 * time.Second)
}
