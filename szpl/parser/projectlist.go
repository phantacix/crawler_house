package parser

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/tebeka/selenium"
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
	pn, _ := pages.Text()
	pageNum, _ := strconv.Atoi(pn)
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

	//a, err := webDriver.FindElements(selenium.ByXPATH, "//a")
	//if err != nil{
	//	panic(err)
	//}
	//result := engine.ParseResult{}
	//
	////  下面这段用于减少 seeds， 只爬一个项目
	//fmt.Printf("%v", a)
	//result.Items = append(result.Items, "华润城润府三期")
	//result.Requests = append(result.Requests, engine.Request{
	//	Url: string("http://ris.szpl.gov.cn/bol/projectdetail.aspx?id=34432"), // 调试
	//	ParserFunc: ParseProject,
	//})

	result := engine.ParseResult{}
	for i := 0; i < pageNum - 106; i++{
		wd, err := webDriver.FindElement(selenium.ByXPATH, "//*[@id=\"AspNetPager1\"]/div[2]/a[3]")
		if err != nil{
			panic(err)
		}
		wd.Click()
		webDriver.CurrentWindowHandle()
		a, err := webDriver.FindElements(selenium.ByXPATH, "//a")
		//if err != nil{
		//	panic(err)
		//}
		for _, v := range a {
			href, _ := v.GetAttribute("href")
			p, _ := v.Text()
			if p != "" {
				// 匹配 url 中的 projectdetail 过滤掉许可证详细信息页面
				// http://ris.szpl.gov.cn/bol/projectdetail.aspx?id=34792
				// http://ris.szpl.gov.cn/bol/certdetail.aspx?id=35373
				if strings.Contains(href, "projectdetail") {
					fmt.Printf("超链接地址：%s, 项目：%s \n", href, p)
					result.Items = append(result.Items, p)
					result.Requests = append(result.Requests, engine.Request{
						Url: string(href),
						ParserFunc: ParseProject,
					})
				}
			}
		}
	}

	//for _, v := range a {
	//	href, _ := v.GetAttribute("href")
	//	p, _ := v.Text()
	//	if p != "" {
	//		// 匹配 url 中的 projectdetail 过滤掉许可证详细信息页面
	//		// http://ris.szpl.gov.cn/bol/projectdetail.aspx?id=34792
	//		// http://ris.szpl.gov.cn/bol/certdetail.aspx?id=35373
	//		if strings.Contains(href, "projectdetail") {
	//			fmt.Printf("超链接地址：%s, 项目：%s \n", href, p)
	//			result.Items = append(result.Items, p)
	//			result.Requests = append(result.Requests, engine.Request{
	//				Url: string(href),
	//				//Url: string("http://ris.szpl.gov.cn/bol/projectdetail.aspx?id=34432"), // 调试
	//				//ParserFunc: engine.NilParser,   // 调试
	//				ParserFunc: ParseProject,
	//			})
	//		}
	//	}
	//}

	return result
}


