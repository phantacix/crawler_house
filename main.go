package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	resp, err := http.Get("http://ris.szpl.gov.cn/bol/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Errror: status code", resp.StatusCode)
		return
	}
	index, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", index)
}
