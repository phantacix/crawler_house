package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, err) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}

	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	// 自动推断：此例效果不好
	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)

}