package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url string

func init() {
	flag.StringVar(&url, "u", "", "input url")
}

func main() {
	// var url = "https://v2ex.com"
	// var errUrl = "https://vvvv.v2.co"
	// if flag.NArg() < 1 {
	// 	panic("参数错误")
	// }
	flag.Parse()
	res, err := http.Get(url)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	checkError(err)
	fmt.Print(string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("er: %v", err)
	}
}
