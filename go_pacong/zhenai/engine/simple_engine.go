package engine

import (
	"log"
	"pacong/fetcher"
)


type SimpleEngine struct { }

func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _, value := range seeds {
		requests = append(requests,value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		//log.Printf("Fetching %s", r.Url)
		//
		//body ,err := fetcher.Fetch(r.Url)
		//if err != nil{
		//	log.Printf("fetcher:err %s :%v", r.Url,err)
		//	continue
		//}
		//
		//parseResult := r.ParserFunc(body)
		parseResult,err := worker(r)
		if err != nil {
			continue
		}


		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items{

			log.Printf("Got item %v",item)
		}
	}
}


func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching URL: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil

}