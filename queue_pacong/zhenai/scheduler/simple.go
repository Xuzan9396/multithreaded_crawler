package scheduler

import (
	"pacong/zhenai/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}


func (s *SimpleScheduler) Submit(r engine.Request)  {

	//fmt.Println("rrr", r);

	go func() {
		s.workerChan <- r
	}()
	//
}


