package main

import (
	"pacong/zhenai/engine"
	"pacong/zhenai/parser"
	"pacong/zhenai/scheduler"
)

//golang.org/x/text
// golang.org/x/net/html
func main() {

	// 多任务版本怕从
	engine.ConcuurentEngineEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 单任务版本
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//printWeiboList(all)
	//fmt.Printf("%s\n",all)
}

//
//func printWeiboList(contents []byte)  {
//	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+">[^<]+</a>`)
//	//<a href="/weibo?q=%23%E9%AD%8F%E5%A4%A7%E5%8B%8B%E7%9C%8B%E9%BB%84%E5%9C%A3%E4%BE%9D%E7%BB%84%E7%9A%84%E8%A1%A8%E6%83%85%23&Refer=top" target="_blank">魏大勋看黄圣依组的表情</a>
//	//re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
//	re := regexp.MustCompile(`<a href="(/weibo\?q=.+)" [^>]+>([^<]+)</a>`)
//	byptes := re.FindAllSubmatch(contents,-1)
//
//	for key, value := range byptes {
//		fmt.Println("key:", key);
//		for k, v:= range value {
//			if k > 0 {
//				if k == 1 {
//					fmt.Printf("https://s.weibo.com%s\n",v);
//
//				}else {
//					fmt.Printf("%s\n",v);
//
//				}
//			}
//		}
//		fmt.Println();
//	}
//	fmt.Printf("match num:%d \n",len(byptes))
//}
