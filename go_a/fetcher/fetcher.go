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
	"strings"
	"time"
)

func Fetch(url string )(respBypte []byte ,err error)  {
	resp,err := http.Get(url)
	//resp,err := http.Get("https://s.weibo.com/top/summary")
	if err != nil{
		return
	}

	defer resp.Body.Close()

	//maps := make(map[string]string)
	//maps["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36"
	//resp,err := HttpPostGet(url,"GET","", maps)
	//if err != nil{
	//	fmt.Errorf("get url 错误 %v", err)
	//	return
	//}


	if resp.StatusCode != http.StatusOK {

		err = fmt.Errorf("wrong status code: %d",resp.StatusCode)
		return
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)


	utf8Reader := transform.NewReader(
		bodyReader,e.NewDecoder(),
	)
	//return utf8Reader;
	return  ioutil.ReadAll(utf8Reader)

}



func FetchIo(url string )(respBypte []byte ,err error)  {

	maps := make(map[string]string)
	maps["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.122 Safari/537.36"
	resp,err := HttpPostGet(url,"GET","", maps)
	if err != nil{
		fmt.Errorf("get url 错误 %v", err)
		return
	}
	return resp,err


	//bodyReader := bufio.NewReader(strings.NewReader(string(resp)))
	//bodyReader := bufio.NewReader(bytes.NewReader(resp))
	//e := determineEncoding(bodyReader)
	//
	//
	//utf8Reader := transform.NewReader(
	//	bodyReader,e.NewDecoder(),
	//)
	//return  ioutil.ReadAll(utf8Reader)
	//return resp,nil


}


func HttpPostGet(url, method, params string, headerArr ...map[string]string) (bytes []byte,err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	request, err := http.NewRequest(method, url, strings.NewReader(params))

	// 设置header

	if method == "POST" && len(headerArr) == 0{
		headerArr[0]["Content-Type"] = "application/json;charset=utf-8"
		//request.Header.Add("Content-Type", "application/json;charset=utf-8")
		//request.Header.Add("Content-Type", "application/json;charset=utf-8")
	}

	if len(headerArr) > 0 {
		for key, value := range headerArr[0] {
			//req.Header.Add("Authorization", "3eex8dY04BdU1amui6bf20ECgtyc9s")
			request.Header.Add(key, value)
		}
	}


	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("wrong status code: %d",response.StatusCode)
		return
	}

	//bodys = response.Body
	return ioutil.ReadAll(response.Body)
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