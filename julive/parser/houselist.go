package parser

import (
	"crawler/engine"
	"regexp"
)

const houseListRe = `<a class="pic-image" href="https://sh.julive.com/project/([0-9]+).html".*alt="([^"]+)`

func ParserArea(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(houseListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://sh.julive.com/project/" + string(m[1]) + "/details.html",
			ParserFunc: ParserProfile,
		})
		result.Items = append(result.Items, "楼盘："+string(m[2]))
	}
	return result
}
