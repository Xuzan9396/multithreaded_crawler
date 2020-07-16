package scheduler

import (
	"pacong/zhenai/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	//panic("implement me")
	return make(chan engine.Request)
}



func (s *QueuedScheduler) Submit(r engine.Request)  {
	s.requestChan <- r // 第二步
}

// 存入worker
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w  // 第一步  10 个worker
}

func (s *QueuedScheduler)Run()  {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ  []engine.Request
		var workerQ []chan engine.Request // 一开始10个队列
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0{ // 第5步
				activeRequest = requestQ[0]
				activeWorker  = workerQ[0] //chan 直接是引用传递 ,此时activeWorker 就是通道worker
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ,r)
			case w := <-s.workerChan:  // 第1步
 				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest: // 第6步
				requestQ = requestQ[1:]
				//fmt.Println(activeWorker,111111);
				workerQ = workerQ[1:]
			}
		}
	}()
}


