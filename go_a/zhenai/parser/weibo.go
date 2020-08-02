package parser

import (
	"fmt"
	"pacong/zhenai/engine"
	"regexp"
	"strings"
)

func WeiboList(contents []byte, url string) engine.ParseResult {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+">[^<]+</a>`)
	//<a href="/weibo?q=%23%E9%AD%8F%E5%A4%A7%E5%8B%8B%E7%9C%8B%E9%BB%84%E5%9C%A3%E4%BE%9D%E7%BB%84%E7%9A%84%E8%A1%A8%E6%83%85%23&Refer=top" target="_blank">魏大勋看黄圣依组的表情</a>
	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)

	result := engine.ParseResult{}
	re := regexp.MustCompile(`<a href="(/weibo\?q=.+)" [^>]+>([^<]+)</a>`)
	byptes := re.FindAllSubmatch(contents, -1)

	var strArr = []string{}

	for key, value := range byptes {
		//fmt.Println("key:", key);
		//var str string

		//str := `%d ：<a href="https://s.weibo.com%s">%s</a>`

		str := `%d ：[%s](https://s.weibo.com%s)`
		//str := `%d ：%s`

		var key1, key2 string
		for k, v := range value {
			if k > 0 {
				if k == 1 {
					key1 = string(v)
					//key1 = ""
					//str = fmt.Sprintf(str,string(v),"%s")

					//fmt.Println(str);
					//str += fmt.Sprintf("%d: https://s.weibo.com%s\n",key,v);

				} else {
					key2 = string(v)
					//str =fmt.Sprintf(str,string(v));

					//fmt.Println(str);
				}
			}

		}
		str = fmt.Sprintf(str, key, key2, key1)

		strArr = append(strArr, str)
		//if key == 3 {
		//	break
		//}
	}
	str := strings.Join(strArr, "\n")

	str = "```\n" + str + "```"

	str = fmt.Sprintf("desp=%s&text=%s", str, "微博热榜")

	result.Items = append(result.Items, str)
	result.Requests = append(result.Requests, engine.Request{
		Url:        url,
		ParserFunc: engine.NilParser,

		//ParserFunc: func(bytes []byte) engine.ParseResult {
		//	return WeiboList(bytes,url)
		//},
	})
	return result
}
