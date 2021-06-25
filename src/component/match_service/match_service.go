package match_service

import (
	"bytes"
	"container/list"
	"demo_go/net_demo/service/util"
	"fmt"
	"go.uber.org/atomic"
	"time"
)

const (
	Insert = 1
	Remove = 2
	Match  = 3

	FinishedType_MaxNode = 1
	FinishedType_OutTime = 2

	Match_OutTime = 3 //匹配超时时间 3秒
	MinMatchSize  = 2 //超时匹配最小匹配大小
)

// 完成匹配回调方法
type OverMatchCallback func(finishedType int, linkedlist *list.List, attach interface{})

// 结点状态回调
type NodeStateCallback func(nodeId int64, state int) // 状态回调

type MatchAttach struct {
	Rid int32 //策划表id
}

// 匹配结点
type MatchNode struct {
	NodeId   int64             // 结点id
	Callback NodeStateCallback // 状态回调
}

type MatchService struct {
	isRuning *atomic.Bool //服务状态

	linkedlist *list.List
	nodeMap    map[int64]*list.Element

	insertChan  chan *MatchNode
	removeChan  chan int64
	commandChan chan string
	ticker      *time.Ticker //ticker驱动器

	flushTime int64 //刷新時間

	maxNode           int               // 最大节点数
	overMatchCallback OverMatchCallback // 匹配回调
	attach            interface{}       // 附件
}

func NewMatchService(maxNode int, attach interface{}, overMatchCallback OverMatchCallback) *MatchService {
	ms := &MatchService{
		isRuning: atomic.NewBool(false),

		linkedlist:        list.New(),
		nodeMap:           map[int64]*list.Element{},
		insertChan:        make(chan *MatchNode, 256),
		removeChan:        make(chan int64, 256),
		commandChan:       make(chan string),
		ticker:            time.NewTicker(1 * time.Second), //1秒驱动一次
		maxNode:           maxNode,
		overMatchCallback: overMatchCallback,
		attach:            attach,
	}
	return ms
}

func (s *MatchService) Startup() {
	isSuccess := s.isRuning.CAS(false, true)
	if isSuccess {
		go func() {
		matchfor:
			for {
				select {
				case newNode := <-s.insertChan:
					s.insertNode(newNode)
					break
				case deleteNodeId := <-s.removeChan:
					s.removeNode(deleteNodeId)
					break
				case <-s.ticker.C:
					s.tick()
					break
				case command := <-s.commandChan:
					switch command {
					case "show":
						s.printNode()
						break
					case "shutdown":
						close(s.commandChan)
						break matchfor
					}
					break
				}
			}
		}()
	}
}

func (s *MatchService) Shutdown() {
	isSuccess := s.isRuning.CAS(true, false)
	if isSuccess {
		fmt.Println("Shutdown")
		s.ticker.Stop()
		close(s.insertChan)
		close(s.removeChan)
		s.commandChan <- "shutdown"
	}
}

func (s *MatchService) InsertNode(node *MatchNode) {
	if !s.isRuning.Load() {
		return
	}
	s.insertChan <- node
}

func (s *MatchService) RemoveNode(nodeId int64) {
	if !s.isRuning.Load() {
		return
	}
	s.removeChan <- nodeId
}

func (s *MatchService) ShowNode() {
	if !s.isRuning.Load() {
		return
	}
	s.commandChan <- "show"
}

func (s *MatchService) printNode() {
	//顺序遍历
	var buffer bytes.Buffer
	n := s.linkedlist.Front()
	for {
		if n == nil {
			break
		}
		matchNode := n.Value.(*MatchNode)
		buffer.WriteString(util.Int642String(matchNode.NodeId))
		buffer.WriteString(",")
		n = n.Next()
	}
	fmt.Println("show ", buffer.String())
}

func (s *MatchService) reset() {
	s.linkedlist = list.New()
	s.nodeMap = map[int64]*list.Element{}
}

func (s *MatchService) insertNode(node *MatchNode) {
	n, ok := s.nodeMap[node.NodeId]
	if !ok {
		if s.flushTime <= 0 {
			s.flushTime = time.Now().Unix()
		}
		n = s.linkedlist.PushBack(node)
		s.nodeMap[node.NodeId] = n
		node.Callback(node.NodeId, Insert)
		//检查匹配
		s.checkMatch()
	}
}

func (s *MatchService) tick() {
	if !s.isRuning.Load() {
		return
	}
	if s.flushTime <= 0 {
		return
	}
	if time.Now().Unix()-s.flushTime < Match_OutTime {
		return
	}
	if s.linkedlist.Len() < MinMatchSize {
		return
	}
	//匹配上
	s.flushTime = 0
	s.overMatchCallback(FinishedType_OutTime, s.linkedlist, s.attach)
	s.reset()
}

func (s *MatchService) checkMatch() {
	if s.linkedlist.Len() >= s.maxNode {
		//匹配上
		s.flushTime = 0
		s.overMatchCallback(FinishedType_MaxNode, s.linkedlist, s.attach)
		s.reset()
	}
}

func (s *MatchService) removeNode(nodeId int64) {
	n, ok := s.nodeMap[nodeId]
	if ok {
		s.linkedlist.Remove(n)
		delete(s.nodeMap, nodeId)

		node := n.Value.(*MatchNode)
		node.Callback(node.NodeId, Remove)
	}
}
