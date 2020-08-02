package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"pacong/zhenai/engine"
)

type AInfo struct {
	Name          string `json:"name"`
	Num           string `json:"num"`
	Add           string `json:"add"`
	AddPercentage string `json:"add_percentage"`
}

func Aparser(body []byte, url string) engine.ParseResult {
	// Load the HTML document
	result := engine.ParseResult{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return result
	}

	data := make([]AInfo, 0)

	doc.Find("#QBS_2_inner>tbody>.LeftLiContainer").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		infoData := AInfo{}
		s.Find("td").Each(func(s int, selection *goquery.Selection) {
			band := selection.Text()
			switch s {
			case 1:
				infoData.Name = band
			case 2:
				infoData.Num = band
			case 3:
				infoData.Add = band
			case 4:
				infoData.AddPercentage = band

			}

		})

		data = append(data, infoData)
	})

	str := ""
	for key, value := range data {
		//fmt.Printf("%d - %-10s - %-10s - %-10s - %-s \n",key,value.Name,value.Num,value.Add,value.AddPercentage)
		str += fmt.Sprintf("%d  %s  %s  %s  %s \n", key, value.Name, value.Num, value.Add, value.AddPercentage)
	}

	str = "```\n" + str + "```"
	str = fmt.Sprintf("desp=%s&text=%s", str, "a股行情")

	//fmt.Println(str);
	result.Items = append(result.Items, str)

	result.Requests = append(result.Requests, engine.Request{
		Url:        url,
		ParserFunc: engine.NilParser,
		//ParserFunc: func(i []byte) engine.ParseResult {
		//	return Aparser(i,url)
		//},

	})
	return result
}
