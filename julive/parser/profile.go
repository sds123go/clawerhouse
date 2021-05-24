package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
)

const StrRe = `<div class="th">楼盘名称</div>[\s\S]<div class="td">[\s\S]<div class="txt">(.*)</div>`

func ParserProfile(content []byte) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(StrRe)
	match := re.FindSubmatch(content)
	fmt.Println("$$$$$$$$$",match)
	if match != nil {
		fmt.Println(string(match[1]))
		profile.BasicIfo.Name = string(match[1])

	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
