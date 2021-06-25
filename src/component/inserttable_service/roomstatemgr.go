package inserttable_service

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const (
	KeyTemplate = "rid_%d"
)

type RoomStateMgr struct {
	pool *redis.Pool
}

func NewRoomStateMgr(pool *redis.Pool) *RoomStateMgr {
	roomStateMgr := &RoomStateMgr{
		pool: pool,
	}
	return roomStateMgr
}

func (s *RoomStateMgr) getKey(rid int32) string {
	return fmt.Sprintf(KeyTemplate, rid)
}

func (s *RoomStateMgr) saveRoomState(roomState *RoomState) bool {
	conn := s.pool.Get()
	defer conn.Close()
	roomStateStr, err := json.Marshal(roomState)
	if err != nil {
		return false
	}
	key := s.getKey(roomState.Rid)
	r, err := redis.Int(conn.Do("HSET", key, roomState.RoomId, roomStateStr))
	if err != nil {
		return false
	}
	return r > 0
}

func (s *RoomStateMgr) removeRoomState(rid int32, roomId int32) bool {
	conn := s.pool.Get()
	defer conn.Close()
	key := s.getKey(rid)
	r, err := redis.Int(conn.Do("HDEL", key, roomId))
	if err != nil {
		return false
	}
	return r > 0
}

func (s *RoomStateMgr) getAllRoomState(rid int32) []*RoomState {
	conn := s.pool.Get()
	defer conn.Close()
	key := s.getKey(rid)
	roomStateStrs, err := redis.Strings(conn.Do("HVALS", key))
	roomStates := make([]*RoomState, len(roomStateStrs))
	if err != nil {
		return roomStates
	}
	for i, v := range roomStateStrs {
		rs := &RoomState{}
		json.Unmarshal([]byte(v), rs)
		roomStates[i] = rs
	}
	return roomStates
}
