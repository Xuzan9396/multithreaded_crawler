package parser

import (
	"pacong/zhenai/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{

	re := regexp.MustCompile(cityListRe)
	byptes := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
	for _, value := range byptes {
		result.Items = append(result.Items,string(value[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(value[1]),
			ParserFunc:engine.NilParser,  // 解析城市
		})

	}
	return result
}


