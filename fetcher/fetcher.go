package fetcher

import (
	"fmt"
	"os"
	"github.com/tebeka/selenium"
	"log"
	"github.com/tebeka/selenium/chrome"
)


func Fetch(url string)(driver selenium.WebDriver) {
	const CHROME_DRIVER_PATH = "C:/Program Files (x86)/Google/Chrome/Application/chromedriver.exe"

	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.

	}
	caps := selenium.Capabilities{"browserName": "chrome"}


	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--disable-gpu",
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)

	selenium.SetDebug(false)
	_, err := selenium.NewChromeDriverService(CHROME_DRIVER_PATH, 9515, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//defer service.Stop()   // 这里不能关闭

	// Connect to the WebDriver instance running locally.
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		log.Printf("Error Connect to the WebDriver instance: %v", err)
	}
	//defer webDriver.Quit()  // 在 parser 解析完毕后再关闭

	//fmt.Println(reflect.TypeOf(webDriver)) //*selenium.remoteWD
	webDriver.Get(url)
	return webDriver
}
