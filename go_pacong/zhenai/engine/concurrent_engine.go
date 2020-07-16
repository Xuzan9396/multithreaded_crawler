package engine

import (
	"log"
	"pacong/fetcher"
)


type ConcuurentEngineEngine struct {
	Scheduler Scheduler
	WorkerCount int

}

type Scheduler interface {
	Submit(Request)
	ConfiguerMasterWorkerChan(chan Request)
}

func (e ConcuurentEngineEngine)Run(seeds ...Request)  {


	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfiguerMasterWorkerChan(in) // 初始化in

	for i := 0; i < e.WorkerCount ; i++  {
		e.createWorker(in,out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for{
		result := <- out
		for _, item:= range result.Items {
			log.Printf("Got item : %v", item)

		}
		for _,request := range result.Requests {
			e.Scheduler.Submit(request) //  超过 10个request ,卡主了， 开启了10个worker, 发了10个 继续发，卡住了,所以需要释放out 不阻塞out
		}
	}


}

func (e ConcuurentEngineEngine)createWorker(in chan Request,out chan ParseResult)  {
	go func() {
		for{
			request := <- in  // 收到10个
			result,err := e.worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}


func (e *ConcuurentEngineEngine)worker(r Request) (ParseResult, error) {
	log.Printf("Fetching URL: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil

}