package parser

import (
	"pacong/zhenai/engine"
	"regexp"
)

//const userRegExp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRegExp = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`
//
//var userRegExpCom = regexp.MustCompile(userRegExp)
//var cityUrlRegExpCom = regexp.MustCompile(cityUrlRegExp)
//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
var(
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)
)


func ParseCity(contents []byte) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	byptes := cityRe.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}

	for _, value := range byptes {
		result.Items = append(result.Items,string(value[2]))

		result.Requests = append(result.Requests,engine.Request{
			//Url:string(value[1]), // 加密了解密不了

			Url:string("https://www.baidu.com"),
			//ParserFunc: func(bytes []byte) engine.ParseResult{
			//	return PareProfile(bytes,name)
			//}, // 解析用户详情
			ParserFunc: func (str string) func(bytes []byte) engine.ParseResult {
				return func(bytes []byte) engine.ParseResult {
					return engine.NilParser(bytes)
					//return PareProfile(bytes,str)
				}
			}(string(value[2])),
		})

	}

	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, value := range matches {
		//fmt.Println(2222222, string(value[1]));
		result.Requests = append(result.Requests,engine.Request{
			Url:string(value[1]),
			ParserFunc:ParseCity,
		})
	}

	return result
}
