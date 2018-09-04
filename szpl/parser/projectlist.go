package parser

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
	"../../engine"
)

func ParseProjectList(driver selenium.WebDriver) engine.ParseResult {

	webDriver := driver
	defer webDriver.Quit()
	// 获取总页数
	pages, err := webDriver.FindElement(selenium.ByXPATH, "//*[@id=\"AspNetPager1\"]/div[1]/b[3]")
	if err != nil {
		panic(err)
	}
	pageNum, _ := pages.Text()
	fmt.Println("Total pages: ", pageNum) //Total pages:  106

	// 获取总项目数
	projectsTotal, err := webDriver.FindElement(selenium.ByXPATH, "//*[@id=\"AspNetPager1\"]/div[1]/b[2]")
	if err != nil {
		panic(err)
	}
	projectNum, _ := projectsTotal.Text()
	fmt.Println("Total projects: ", projectNum) // Total projects:  3169

	// 获取项目名称
	projects, err := webDriver.FindElement(selenium.ByID, "DataList1")
	if err != nil {
		panic(err)
	}
	project, _ := projects.Text()
	fmt.Println("Total projects: ", project)

	a, err := webDriver.FindElements(selenium.ByXPATH, "//a")
	if err != nil{
		panic(err)
	}
	result := engine.ParseResult{}
	for _, v := range a {
		href, _ := v.GetAttribute("href")
		p, _ := v.Text()
		if p != "" {
			if strings.HasPrefix(href, "http://") {
				fmt.Printf("超链接地址：%s, 项目：%s \n", href, p)
				result.Items = append(result.Items, p)
				result.Requests = append(result.Requests, engine.Request{
					Url: string(href),
					ParserFunc: engine.NilParser,
				})
			}
		}
	}
	return result
}


