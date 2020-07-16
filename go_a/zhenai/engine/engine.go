package engine

import (
	"fmt"
	"log"
	"pacong/fetcher"
	"time"
)

func Run(seeds ...Request)  {
	var requests []Request
	var notify chan int
	for _, value := range seeds {
		requests = append(requests,value)
	}


	for {
		lens := len(requests)
		fmt.Println("lens:", lens);
		if len(requests) > 0  {
			r := requests[0]
			requests = requests[1:]
			log.Printf("Fetching %s", r.Url)

			body ,err := fetcher.FetchIo(r.Url)
			if err != nil{
				log.Printf("fetcher:err %s :%v", r.Url,err)
				continue
			}

			parseResult := r.ParserFunc(body)
			requests = append(requests,parseResult.Requests...)
			for _,item := range parseResult.Items {
				log.Printf("Got item %v",item)

			}

		}else{
			time.Sleep(1 * time.Second)
		}



		//time.Sleep(5 * time.Second)




		//parseResult := r.ParserFunc(res.body)
		//requests = append(requests,parseResult.Requests...)
		//for _,item := range parseResult.Items{
		//
		//	log.Printf("Got item %v",item)
		//}
	}


	<- notify

}


