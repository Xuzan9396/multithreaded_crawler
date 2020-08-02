package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"pacong/fetcher"
	"regexp"
	"testing"
)

func TestParseProfileList(t *testing.T) {

	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1028370723")
	//contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := PareProfile(contents, "xuzan")
	fmt.Println(result)

	//const resultSize = 470
	//if len(result.Requests) != resultSize {
	//	t.Errorf("result %d \n", len(result.Requests) )
	//}else{
	//	for _, value := range result.Requests {
	//		fmt.Println(value.Url);
	//	}
	//
	//	fmt.Printf("%s\n",result.Items)
	//	fmt.Println(len(result.Items));
	//}

}

func TestParseProfileList2(t *testing.T) {

	url := "https://album.zhenai.com/u/1028370723"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "FSSBBIl1UgzbN7NO=5T6.6rur2xZYEln9TliryX_KMpQTBsYgo4sxzlMNBBIXTGJpb7q3g8uYzhXBuUhcp43wF3xa5ehL1O7pBg7FvwA")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func TestParseProfileList3(t *testing.T) {
	content := "wxid_5953469534512:\n测试2"
	var ageRe = regexp.MustCompile(`wxid_[0-9a-z]+:\n(.+)`)

	byptes := ageRe.FindStringSubmatch(content)
	for _, value := range byptes {
		fmt.Println(value)

	}

}
