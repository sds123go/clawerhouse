package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strings"
)

var NameRe = regexp.MustCompile(`<div class="th">楼盘名称</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">([^<]*)`)
var PriceRe = regexp.MustCompile(`<div class="th">参考单价</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">([^<]*)`)
var TotalRe = regexp.MustCompile(`.*总价.*</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">([^<]*)`)
var PropertyRe = regexp.MustCompile(`<div class="th">物业类型</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">([^<]*)`)
var AddressRe = regexp.MustCompile(`<div class="th">楼盘地址</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt add-txt"">([^<]*)`)
var LoopRe = regexp.MustCompile(`<div class="th">环&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;线</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">([^<]*)`)
var OpenRe = regexp.MustCompile(`<div class="th">开盘时间</div>[\s\S]+?<div class="td .*">[\s\S]+?<div class="txt">([^<]*)`)
var HandRe = regexp.MustCompile(`<div class="th">.*交房.*</div>[\s\S]+?<div class="td .*">[\s\S]+?<div class="txt">([^<]*)`)
var DeveloperRe = regexp.MustCompile(`<div class="th">开 发 商</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">[\s\S]+?<div class="name .*" title="([^"]*)`)
var GreeneryRe = regexp.MustCompile(`<div class="th">绿 化 率</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt .*">([^<]*)`)
var BuildTypeRe = regexp.MustCompile(`<div class="th">建筑类型</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt .*">([^<]*)`)
var PeriodIntRe = regexp.MustCompile(`<div class="th">产权年限</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt .*">([^<]*)`)
var PlotRatioRe = regexp.MustCompile(`<div class="th">容 积 率</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt .*">([^<]*)`)
var CompanyRe = regexp.MustCompile(`<div class="th">物业公司</div>[\s\S]+?<div class="td">[\s\S]+?<div class="txt">[\s\S]+?<div class=".*" title="([^"]*)`)
var PropertyPriceRe = regexp.MustCompile(`<div class="th">物业费用</div>[\s\S]+?<div class="td">[\s\S]+?<div class=".*">([^<]*)`)
var WaterGasRe = regexp.MustCompile(`<div class="th">水电燃气</div>[\s\S]+?<div class="td">[\s\S]+?<div class=".*">([^<]*)`)

func ParserProfile(content []byte) engine.ParseResult {
	profile := model.Profile{}
	profile.BasicIfo.Name = extractResult(content, NameRe)
	profile.BasicIfo.AvgPrice = extractResult(content, PriceRe)
	profile.BasicIfo.TotalPrice = strings.TrimSpace(extractResult(content, TotalRe))
	profile.BasicIfo.PropertyType = extractResult(content, PropertyRe)
	profile.BasicIfo.Address = extractResult(content, AddressRe)
	profile.BasicIfo.Loop = extractResult(content, LoopRe)
	profile.SaleIfo.OpenTime = extractResult(content, OpenRe)
	profile.SaleIfo.HandTime = extractResult(content, HandRe)
	profile.BuildIfo.Developer = extractResult(content, DeveloperRe)
	profile.BuildIfo.Greenery = extractResult(content, GreeneryRe)
	profile.BuildIfo.BuildType = extractResult(content, BuildTypeRe)
	profile.BuildIfo.PeriodInt = extractResult(content, PeriodIntRe)
	profile.BuildIfo.PlotRatio = extractResult(content, PlotRatioRe)
	profile.PropertyIfo.Company = extractResult(content, CompanyRe)
	profile.PropertyIfo.PropertyPrice = extractResult(content, PropertyPriceRe)
	profile.PropertyIfo.WaterGas = extractResult(content, WaterGasRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
func extractResult(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
