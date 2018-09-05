package parser

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
	"../../engine"
)

func ParseBranchList(driver selenium.WebDriver) engine.ParseResult {

	webDriver := driver
	defer webDriver.Quit()


	a, err := webDriver.FindElements(selenium.ByXPATH, "//a")
	if err != nil{
		panic(err)
	}
	result := engine.ParseResult{}
	for _, v := range a {
		href, _ := v.GetAttribute("href")
		p, _ := v.Text()
		if p != "" {
			// 匹配 url 中的 hezuo  过滤掉合作方信息
			// http://ris.szpl.gov.cn/bol/building.aspx?id=31784&presellid=35012
			// http://ris.szpl.gov.cn/bol/hezuo.aspx?id=35012
			if strings.Contains(href, "Branch") {
				fmt.Printf("超链接地址：%s, building：%s \n", href, p)
				result.Items = append(result.Items, p)
				result.Requests = append(result.Requests, engine.Request{
					Url: string(href),
					//ParserFunc: engine.NilParser,  // 调试
					ParserFunc: ParseHouseList,
				})
			}
		}
	}
	return result
}


