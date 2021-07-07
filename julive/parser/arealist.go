package parser

import (
	"crawler/engine"
	"regexp"
)

const areaListRe = `<a class="" href="(https://sh.julive.com/project/s/([0-9a-z]+))"`

//const nextpageRe = `<li><a href="(https://sh.julive.com/project/s/.*)" target="_self" data-page=".">.</a></li>`
func ParseAreaList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(areaListRe)
	matches := re.FindAllSubmatch(contents, 16)
	//nextre:=regexp.MustCompile(nextpageRe)
	//NextPage:=nextre.FindSubmatch(contents)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserArea,
		})
		result.Items = append(result.Items, engine.Item{
			Url:     string(m[1]),
			Type:    "area",
			Id:      string(m[2]),
			Payload: "行政区:" + string(m[2]),
		})
	}
	// for _,n:=range NextPage{
	// 	result.Request=append(result.Request, engine.Request{
	// 		Url: string(n[1]),
	// 		ParserFunc: ParserArea,
	// 	})
	// }
	return result
}
