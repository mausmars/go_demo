package product_customer_service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
生产消费模式測試
*/
type WorkerFactory struct {
	id int
}

func (s *WorkerFactory) createWorker(id int) IWorker {
	worker := &Worker{id: id}
	return worker.work
}

//-----------------------------
type Worker struct {
	id int
}

func (s *Worker) work(task Task) {
	fmt.Println("work_id=", s.id, " taks=", task)
	s.randomSleep(30)
}

func (s *Worker) randomSleep(num int) {
	time.Sleep(time.Duration(rand.Intn(num)) * time.Millisecond)
}

//-----------------------------
func TestPCService(t *testing.T) {
	factory := &WorkerFactory{}
	s := NewPCService(24, factory.createWorker, 10)
	s.Startup()
	go func() {
		for i := 0; i < 100000; i++ {
			s.Product(i)
		}
		fmt.Println("Product Over!")
	}()
	//time.Sleep(10 * time.Second)
	//s.Shutdown()
	time.Sleep(60 * time.Second)
	fmt.Println("Test Over!")
}
