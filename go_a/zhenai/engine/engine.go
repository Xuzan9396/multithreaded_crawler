package engine

import (
	"log"
	"pacong/fetcher"
	"time"
)

var G_RequestChan chan Request

func Run(seeds ...Request) {
	var requests []Request
	var notify chan int
	for _, value := range seeds {
		requests = append(requests, value)
	}
	G_RequestChan = make(chan Request, 10)

	for {
		//lens := len(requests)
		//fmt.Println("lens:", lens);
		if len(requests) > 0 {
			r := requests[0]
			requests = requests[1:]
			//log.Printf("Fetching %s", r.Url)

			body, err := fetcher.FetchIo(r.Url)
			if err != nil {
				log.Printf("fetcher:err %s :%v", r.Url, err)
				continue
			}

			parseResult := r.ParserFunc(body)
			requests = append(requests, parseResult.Requests...)
			for _, item := range parseResult.Items {
				log.Printf("Got item %v", item)
				maps := make(map[string]string)
				maps["Content-Type"] = "application/x-www-form-urlencoded"
				respFangtang, _ := fetcher.HttpPostGet("https://sc.ftqq.com/SCU64514T1d2bceaaf7386be63d2e6da3d22e46995da98d09d8ca7.send", "POST", item.(string), maps)
				log.Println(string(respFangtang))
			}

		}

		select {
		case res := <-G_RequestChan:
			requests = append(requests, res)
		default:
			time.Sleep(time.Second)

		}

	}

	<-notify

}

func SetRequestChan(res Request) {
	if G_RequestChan == nil {
		G_RequestChan = make(chan Request, 10)
	}
	G_RequestChan <- res
}
