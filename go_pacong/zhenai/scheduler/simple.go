package scheduler

import (
	"pacong/zhenai/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfiguerMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request)  {

	//fmt.Println("rrr", r);

	go func() {
		s.workerChan <- r
	}()
	//
}


