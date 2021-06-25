package product_customer_service

/**
生产消费模式測試
*/

import (
	"fmt"
	"go.uber.org/atomic"
)

// 任務
type Task interface{}

// 工人
type IWorker func(Task)

// 工人工厂
type IWorkerFactory func(id int) IWorker

type PCService struct {
	isRuning *atomic.Bool //服务状态

	ch      chan Task // 任务管道
	workers []IWorker // 工人
}

func (s *PCService) Startup() {
	isSuccess := s.isRuning.CAS(false, true)
	if isSuccess {
		fmt.Println("PCService Startup!")
		for _, w := range s.workers {
			go func(ch chan Task, worker IWorker) {
				for {
					task := <-ch
					if task != nil {
						worker(task)
					}
				}
			}(s.ch, w)
		}
	}
}

func (s *PCService) Shutdown() {
	isSuccess := s.isRuning.CAS(true, false)
	if isSuccess {
		fmt.Println("PCService Shutdown!")
		close(s.ch)
	}
}

func (s *PCService) Product(task Task) {
	if !s.isRuning.Load() {
		return
	}
	s.ch <- task
}

func NewPCService(workerCount int, wf IWorkerFactory, chanSize int) *PCService {
	if chanSize < 0 {
		chanSize = 0
	}
	ch := make(chan Task, chanSize)
	workers := make([]IWorker, workerCount)
	for id := 0; id < workerCount; id++ {
		workers[id] = wf(id + 1)
	}
	s := &PCService{
		isRuning: atomic.NewBool(false),
		ch:       ch,
		workers:  workers,
	}
	s.workers = workers
	return s
}

//-----------------------------
