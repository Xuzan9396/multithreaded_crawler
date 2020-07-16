package engine

import (
	"log"
	"pacong/fetcher"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _, value := range seeds {
		requests = append(requests,value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)

		body ,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("fetcher:err %s :%v", r.Url,err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items{

			log.Printf("Got item %v",item)
		}
	}
}
