package parser

import (
	"pacong/model"
	"pacong/zhenai/engine"
	"regexp"
	"strconv"
)
//<td width="180"><span class="grayL">年龄：</span>35</td>


var ageRe= regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var marriageRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)

func PareProfile(contents []byte,name string ) engine.ParseResult  {

	profile := model.Profile{}
	profile.Name = name
	age,err := strconv.Atoi(
		("10"))
		//extractString(contents,ageRe))
	if err != nil {
		age = 0
	}
	profile.Age = age

	//profile.Marriage = extractString(contents,marriageRe) // 婚况

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result


}


func extractString(contents []byte, re *regexp.Regexp) string {

	math := re.FindSubmatch(contents)
	if math != nil && len(math) >= 2 {

		return string(math[1])

	}
	return ""
}