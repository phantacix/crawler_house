package engine

import "github.com/tebeka/selenium"

type Request struct {
	Url string
	ParserFunc func(driver selenium.WebDriver) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParser(driver selenium.WebDriver) ParseResult {
	defer driver.Close()
	return ParseResult{}
}