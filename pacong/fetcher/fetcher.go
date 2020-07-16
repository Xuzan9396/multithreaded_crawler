package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string )(respBypte []byte ,err error)  {
	resp,err := http.Get(url)
	//resp,err := http.Get("https://s.weibo.com/top/summary")
	if err != nil{
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		err = fmt.Errorf("wrong status code: %d",resp.StatusCode)
		return
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)

	//utf8Reader := transform.NewReader(
	//	resp.Body,simplifiedchinese.GBK.NewDecoder(),
	//	)

	utf8Reader := transform.NewReader(
		bodyReader,e.NewDecoder(),
	)
	return  ioutil.ReadAll(utf8Reader)

}


// 获取编码
func determineEncoding(r *bufio.Reader) encoding.Encoding  {

	bytes,err := r.Peek(1024)

	if err != nil {
		log.Printf("Fetcher err: %v", err)
		return unicode.UTF8
		//panic(err)
	}

	e,_,_:= charset.DetermineEncoding(bytes,"")
	//fmt.Println(e,name,certain);
	return e

}