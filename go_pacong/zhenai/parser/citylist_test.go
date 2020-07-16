package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}


	result := ParseCityList(contents)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result %d \n", len(result.Requests) )
	}else{
		for _, value := range result.Requests {
			fmt.Println(value.Url);
		}

		fmt.Printf("%s\n",result.Items)
		fmt.Println(len(result.Items));
	}

}