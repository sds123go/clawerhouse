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
		url := "https://sh.julive.com/project/" + string(m[1]) + "/details.html"
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(b []byte) engine.ParseResult {
				return ParserProfile(b, url, string(m[1]))
			},
		})
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Type:    "house",
			Id:      string(m[1]),
			Payload: "楼盘：" + string(m[2]),
		})
	}
	return result
}
