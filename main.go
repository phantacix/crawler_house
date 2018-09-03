package main

import (
	"github.com/tebeka/selenium"
	"fmt"
	"os"
)

func main() {
	const CHROME_DRIVER_PATH  = "C:/Program Files (x86)/Google/Chrome/Application/chromedriver.exe"
	const URL_START = "http://ris.szpl.gov.cn/bol/"

	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	caps := selenium.Capabilities{"browserName": "chrome"}

	selenium.SetDebug(true)
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
	fmt.Println("Total pages: ", pageNum)       //Total pages:  106


	// 获取总项目数
	projects, err := webDriver.FindElement(selenium.ByXPATH, "//*[@id=\"AspNetPager1\"]/div[1]/b[2]")
	if err != nil {
		panic(err)
	}
	projectNum, _ := projects.Text()
	fmt.Println("Total projects: ", projectNum) // Total projects:  3169

}