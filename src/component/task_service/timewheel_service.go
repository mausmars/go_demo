package task_service

import (
	"container/list"
	"go.uber.org/atomic"
	"time"
)

// 任务回调函数
type TaskCallBack func(userData interface{})

// TaskData 回调函数参数类型

// TimeWheel 时间轮
type TimeWheelService struct {
	isRuning *atomic.Bool //服务状态

	ticker   *time.Ticker
	interval time.Duration // 指针每隔多久往前移动一格

	slots []*list.List // 时间轮槽
	// key: 定时器唯一标识 value: 定时器所在的槽, 主要用于删除定时器, 不会出现并发读写，不加锁直接访问
	timer       map[interface{}]int
	currentPos  int              // 当前指针指向哪一个槽
	slotNum     int              // 槽数量
	callback    TaskCallBack     // 定时器回调函数
	insertChan  chan *Task       // 新增任务channel
	removeChan  chan interface{} // 删除任务channel
	stopChannel chan bool        // 停止定时器channel
}

// Task 任务
type Task struct {
	delay    time.Duration // 延迟时间
	circle   int           // 时间轮需要转动几圈
	key      interface{}   // 定时器唯一标识, 用于删除定时器
	userData interface{}   // 回调函数参数
}

// 创建时间轮（第一个参数为tick刻度, 即时间轮多久转动一次；第二个参数为时间轮槽slot数量；第三个参数为回调函数）
func NewTimeWheelService(interval time.Duration, slotNum int, callback TaskCallBack) *TimeWheelService {
	if interval <= 0 || slotNum <= 0 || callback == nil {
		return nil
	}
	s := &TimeWheelService{
		isRuning:    atomic.NewBool(false),
		interval:    interval,
		slots:       make([]*list.List, slotNum),
		timer:       make(map[interface{}]int),
		currentPos:  0,
		callback:    callback,
		slotNum:     slotNum,
		insertChan:  make(chan *Task),
		removeChan:  make(chan interface{}),
		stopChannel: make(chan bool),
	}
	// 初始化槽，每个槽指向一个双向链表
	for i := 0; i < s.slotNum; i++ {
		s.slots[i] = list.New()
	}
	return s
}

//启动服务
func (s *TimeWheelService) Startup() {
	isSuccess := s.isRuning.CAS(false, true)
	if isSuccess {
		s.ticker = time.NewTicker(s.interval)
		go func() {
		Loop:
			for {
				select {
				case <-s.ticker.C:
					s.tickHandler()
					break
				case task := <-s.insertChan:
					s.addTask(task)
					break
				case key := <-s.removeChan:
					s.removeTask(key)
					break
				case <-s.stopChannel:
					s.ticker.Stop()
					break Loop
				}
			}
		}()
	}
}

// 停止服务
func (s *TimeWheelService) Shutdown() {
	isSuccess := s.isRuning.CAS(true, false)
	if isSuccess {
		s.stopChannel <- true
		close(s.removeChan)
		close(s.insertChan)
	}
}

// 添加定时器（第一个参数为延迟时间；第二个参数为定时器唯一标识, 删除定时器需传递此参数；第三个参数为用户自定义数据, 此参数将会传递给回调函数, 类型为interface{}；）
func (s *TimeWheelService) InsertTask(delay time.Duration, key interface{}, userData interface{}) {
	if !s.isRuning.Load() {
		return
	}
	if delay < 0 {
		return
	}
	s.insertChan <- &Task{
		delay:    delay,
		key:      key,
		userData: userData,
	}
}

// 删除任务
func (s *TimeWheelService) RemoveTask(key interface{}) {
	if !s.isRuning.Load() {
		return
	}
	if key == nil {
		return
	}
	s.removeChan <- key
}

func (s *TimeWheelService) tickHandler() {
	l := s.slots[s.currentPos]
	s.scanAndRunTask(l)
	if s.currentPos == s.slotNum-1 {
		s.currentPos = 0
	} else {
		s.currentPos++
	}
}

// 扫描链表中过期定时器, 并执行回调函数
func (s *TimeWheelService) scanAndRunTask(l *list.List) {
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.circle > 0 {
			task.circle--
			e = e.Next()
			continue
		}
		go s.callback(task.userData)
		next := e.Next()
		l.Remove(e)
		if task.key != nil {
			delete(s.timer, task.key)
		}
		e = next
	}
}

// 新增任务到链表中
func (s *TimeWheelService) addTask(task *Task) {
	pos, circle := s.getPositionAndCircle(task.delay)
	task.circle = circle
	s.slots[pos].PushBack(task)
	if task.key != nil {
		s.timer[task.key] = pos
	}
}

// 获取定时器在槽中的位置, 时间轮需要转动的圈数
func (s *TimeWheelService) getPositionAndCircle(d time.Duration) (pos int, circle int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(s.interval.Seconds())
	circle = delaySeconds / intervalSeconds / s.slotNum
	pos = (s.currentPos + delaySeconds/intervalSeconds) % s.slotNum
	return
}

// 从链表中删除任务
func (s *TimeWheelService) removeTask(key interface{}) {
	// 获取定时器所在的槽
	position, ok := s.timer[key]
	if !ok {
		return
	}
	// 获取槽指向的链表
	l := s.slots[position]
	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.key == key {
			delete(s.timer, task.key)
			l.Remove(e)
		}
		e = e.Next()
	}
}
