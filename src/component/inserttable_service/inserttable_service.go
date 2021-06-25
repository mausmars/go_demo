package inserttable_service

import (
	"container/list"
	"fmt"
	"go.uber.org/atomic"
)

const (
	Command_GetSeatId = 1 //用户插桌指令
	Command_AddSeatId = 2 //需要插桌的房間
)

type RoomState struct {
	RoomServiceId string  `json:"roomServiceId"` //房间所在serviceId
	RoomId        int32   `json:"roomId"`        //房间id
	Rid           int32   `json:"rid"`           //策划表对应id
	SeatIds       []int32 `json:"seatIds"`       //策划表对应id
}

//------------------------------------
//用户插桌指令
type GetSeatIdCommand struct {
	fpid     int64                                                    //用戶id
	rid      int32                                                    //希望的策划表Id
	callback func(isSuccess bool, seatId int32, roomState *RoomState) //是否插桌成功回调
}

//需要插桌的房間
type AddSeatIdCommand struct {
	roomServiceId string //房间所在serviceId
	roomId        int32  //房间id
	rid           int32  //希望的策划表Id
	seatId        int32  //空座位
}

//------------------------------------
type Command struct {
	CommandId int         //指令号
	Attach    interface{} //附件
}

//插桌服务
type InsertTableService struct {
	isRuning *atomic.Bool //服务状态

	//{策划表id:RoomState列表}
	roomStateMap map[int32]*list.List
	//{策划表id:RoomState}
	nodeMap map[int32]*list.Element

	commandChan chan *Command //指令管道
	stopChannel chan bool     // 停止服务

	rids         []int32 //策划表id
	roomStateMgr *RoomStateMgr
}

var InsertTable *InsertTableService

func InitInsertTableService(rids []int32, roomStateMgr *RoomStateMgr) *InsertTableService {
	return NewInsertTableService(rids, roomStateMgr)
}

func NewInsertTableService(rids []int32, roomStateMgr *RoomStateMgr) *InsertTableService {
	ms := &InsertTableService{
		isRuning:     atomic.NewBool(false),
		roomStateMap: map[int32]*list.List{},
		nodeMap:      map[int32]*list.Element{},

		commandChan: make(chan *Command),
		stopChannel: make(chan bool),

		rids:         rids,
		roomStateMgr: roomStateMgr,
	}
	return ms
}

func (s *InsertTableService) Startup() {
	isSuccess := s.isRuning.CAS(false, true)
	if isSuccess {
		// 从redis中取出RoomState初始化
		for _, rid := range s.rids {
			s.roomStateMap[rid] = list.New()

			roomStates := s.roomStateMgr.getAllRoomState(rid)
			for _, roomState := range roomStates {
				if len(roomState.SeatIds) <= 0 {
					continue
				}
				s.insertRoomState(roomState)
			}
		}
		go func() {
		Loop:
			for {
				select {
				case command := <-s.commandChan:
					switch command.CommandId {
					case Command_GetSeatId:
						s.getSeatIdCommand(command)
						break
					case Command_AddSeatId:
						s.addSeatIdCommand(command)
						break
					}
				case <-s.stopChannel:
					break Loop
				}
			}
		}()
	}
}

func (s *InsertTableService) Shutdown() {
	isSuccess := s.isRuning.CAS(true, false)
	if isSuccess {
		fmt.Println("Shutdown")
		s.stopChannel <- true
		close(s.commandChan)
	}
}

func (s *InsertTableService) getSeatIdCommand(command *Command) {
	c := command.Attach.(*GetSeatIdCommand)
	isSuccess, seatId, roomState := s.searchSeatId(c.rid)
	if isSuccess {
		// 成功返回回调
		c.callback(isSuccess, seatId, roomState)
	} else {
		//TODO 降級匹配
		c.callback(isSuccess, seatId, roomState)
	}
}

func (s *InsertTableService) searchSeatId(rid int32) (bool, int32, *RoomState) {
	list, ok := s.roomStateMap[rid]
	if !ok {
		return false, 0, nil
	}
	roomStateE := list.Front()
	if roomStateE == nil {
		return false, 0, nil
	}
	roomState := roomStateE.Value.(*RoomState)
	seatId := roomState.SeatIds[0]
	roomState.SeatIds = roomState.SeatIds[1:]
	if len(roomState.SeatIds) <= 0 {
		list.Remove(roomStateE)             //从队列中移除
		delete(s.nodeMap, roomState.RoomId) //从房间映射中移除
		// 从redis中移除 TODO io操作最好异步处理
		s.roomStateMgr.removeRoomState(roomState.Rid, roomState.RoomId)
	} else {
		// 变更数据更新到redis中 TODO io操作最好异步处理
		s.roomStateMgr.saveRoomState(roomState)
	}
	return true, seatId, roomState
}

func (s *InsertTableService) addSeatIdCommand(command *Command) {
	c := command.Attach.(*AddSeatIdCommand)
	roomStateE, ok := s.nodeMap[c.roomId]
	if ok {
		// 如果存在把空位加入到列表中
		roomState := roomStateE.Value.(*RoomState)
		roomState.SeatIds = append(roomState.SeatIds, c.seatId)
	} else {
		// 如果不存放入新的RoomState
		roomState := &RoomState{
			RoomServiceId: c.roomServiceId,
			RoomId:        c.roomId,
			Rid:           c.rid,
			SeatIds:       []int32{},
		}
		roomState.SeatIds = append(roomState.SeatIds, c.seatId)
		list, ok := s.roomStateMap[c.rid]
		if !ok {
			return
		}
		roomStateE := list.PushBack(roomState)
		s.nodeMap[c.roomId] = roomStateE
	}
}

func (s *InsertTableService) insertRoomState(roomState *RoomState) {
	list, ok := s.roomStateMap[roomState.Rid]
	if !ok {
		return
	}
	roomStateE := list.PushBack(roomState)
	s.nodeMap[roomState.RoomId] = roomStateE
}

func (s *InsertTableService) GetSeatIdCommand(c *GetSeatIdCommand) {
	command := &Command{
		CommandId: Command_GetSeatId,
		Attach:    c,
	}
	s.commandChan <- command
}

func (s *InsertTableService) AddSeatIdCommand(c *AddSeatIdCommand) {
	command := &Command{
		CommandId: Command_AddSeatId,
		Attach:    c,
	}
	s.commandChan <- command
}
