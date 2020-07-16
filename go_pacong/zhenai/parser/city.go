package parser

import (
	"pacong/zhenai/engine"
	"regexp"
)

//const userRegExp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
//const cityUrlRegExp = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`
//
//var userRegExpCom = regexp.MustCompile(userRegExp)
//var cityUrlRegExpCom = regexp.MustCompile(cityUrlRegExp)
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	byptes := re.FindAllSubmatch(contents,-1)

	result := engine.ParseResult{}
/*	func(v []byte) func() error {
		return func() error {
			int64Id, _ := strconv.ParseInt(string(v), 10, 64)
			bytes, err := modelunionsGiftLogs(int64Id)
			if err != nil {
				jobLog.Errorf(err, "")
			} else {
				// 推送日志
				G_logSink.appendJobLog(fmt.Sprintf("%d,%s", int64Id, string(bytes)))
				//jobLog.Infof("ids:%d,%s", int64Id, string(bytes))

			}
			return nil
		}
	}(resPop),*/
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
					return PareProfile(bytes,str)
				}
			}(string(value[2])),
		})

	}
	return result
}
