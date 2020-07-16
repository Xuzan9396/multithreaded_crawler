package main

import (
	"fmt"
	"log"
	"time"
)

func main2() {
   model := QueuedScheduler{}
   model.Run()

	for i := 0; i < 10  ; i++  {
		worker := make(chan int)
		go func(i int) {
			for{
				model.WorkerReady(worker) // 去分发一个worker 让request 分发
				request := <- worker  // 收到10个
				time.Sleep(time.Second)
				log.Printf("worker%d: result_ %v",i,request)
			}
		}(i)
	}

	in := 0
	for{
		for i := 0; i < 14  ; i++ {
			model.Submit(in)
		}
		time.Sleep(time.Second*10)

		in++
	}


}

func main()  {
	a := make(chan int,1)
	b := a
	go func() {
		b <-22
	}()

	res := <-b
	fmt.Println(res);
}


func (s *QueuedScheduler) WorkerReady(w chan int) {
	s.workerChan <- w  // 第一步  10 个worker
}

func (s *QueuedScheduler)Run()  {
	s.requestChan = make(chan int,1)
	s.workerChan = make(chan chan int,1)
	go func() {
		var requestQ  []int
		var workerQ []chan int // 一开始10个队列
		for{
			var activeRequest int
			var activeWorker chan int

			//fmt.Println(len(requestQ),len(workerQ));
			if len(requestQ) > 0 && len(workerQ) > 0{ // 第5步
				fmt.Println(len(requestQ),len(workerQ));
				activeRequest = requestQ[0]
				activeWorker  = workerQ[0] //chan 直接是引用传递 ,此时activeWorker 就是通道worker
				//fmt.Println(1111,activeWorker);  //

			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ,r)
			case w := <-s.workerChan:  // 第1步
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest: // 第6步

				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
				//fmt.Println(2222,activeWorker);
			}
		}
	}()
}


type QueuedScheduler struct {
	requestChan chan int
	workerChan chan chan int
}


func (s *QueuedScheduler) Submit(r int)  {
	s.requestChan <- r // 第二步
}

