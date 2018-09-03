package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
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