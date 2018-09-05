package parser

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
	"../../engine"
	"regexp"
)

func ParseHouseList(driver selenium.WebDriver) engine.ParseResult {

	webDriver := driver
	defer webDriver.Quit()


	a, err := webDriver.FindElements(selenium.ByXPATH, "//a")
	if err != nil{
		panic(err)
	}
	result := engine.ParseResult{}
	reg := regexp.MustCompile("id=.*")
	for _, v := range a {
		href, _ := v.GetAttribute("href")
		id := reg.FindAllString(href, -1)
		// 匹配 url 中的 housedetail  过滤
		// http://ris.szpl.gov.cn/bol/housedetail.aspx?id=1578802
		if strings.Contains(href, "housedetail") {
			fmt.Printf("houselist 超链接地址：%s, houseid：%s \n", href, id)
			result.Items = append(result.Items, id)
			result.Requests = append(result.Requests, engine.Request{
				Url: string(href),
				//ParserFunc: engine.NilParser,  // 调试
				ParserFunc: ParseProfile,
			})
		}
	}
	return result
}
