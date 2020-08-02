package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"time"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("./chromedriver83.0.4103.39")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("https://www.helloweba.net/demo/2017/unlock/")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(5 * time.Second)

	el, err := session.FindElement(webdriver.ClassName, "slide-to-unlock-handle")
	//el,err := session.FindElement(webdriver.XPath,"/html/body/div/div/div/h2/a")
	//log.Printf("%v",el)
	el.Click()
	time.Sleep(1 * time.Second)

	//session.Click(webdriver.LeftButton)
	//session.Click(webdriver.LeftButton)
	//if err := session.MoveTo(el,5,0);err != nil{
	//	fmt.Println(err,"错误");
	//}else{
	//	fmt.Println("移动成功");
	//}
	//el.Click()
	//webdriver.MoveTo

	time.Sleep(60 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
