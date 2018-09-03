package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"strings"
)

func main() {
	const CHROME_DRIVER_PATH = "C:/Program Files (x86)/Google/Chrome/Application/chromedriver.exe"
	const URL_START = "http://ris.szpl.gov.cn/bol/"

	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	caps := selenium.Capabilities{"browserName": "chrome"}

	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(CHROME_DRIVER_PATH, 9515, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer webDriver.Quit()

	// 访问目标网站首页
	err = webDriver.Get(URL_START)
	if err != nil {
		panic(err)
	}

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
	for _, v := range a {
		href, _ := v.GetAttribute("href")
		p, _ := v.Text()
		if p != "" {
			if strings.HasPrefix(href, "http://") {
				fmt.Printf("超链接地址：%s, 项目：%s \n", href, p)
			}
		}
	}
}
