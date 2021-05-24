package parser

import (
	"crawler/engine"
	"regexp"
)

const areaListRe = `<a class="" href="(https://sh.julive.com/project/s/([0-9a-z]+))"`

func ParseAreaList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(areaListRe)
	matches := re.FindAllSubmatch(contents, 16)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserArea,
		})
		result.Items = append(result.Items, string(m[2]))
	}
	return result
}
