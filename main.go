package main

import (
	"./engine"
	"./szpl/parser"
)

func main() {
	const CHROME_DRIVER_PATH = "C:/Program Files (x86)/Google/Chrome/Application/chromedriver.exe"
	const URL_START = "http://ris.szpl.gov.cn/bol/"

	engine.Run(engine.Request{
		Url: URL_START,
		ParserFunc :parser.ParseProjectList,
	})
}
