package main

import "net/http"

import "fmt"

func main() {
	pollUrls := []string{
		"https://google.cn",
		"https://v2ex.com",
		"https://www.baidu.com",
	}
	for _, url := range pollUrls {
		req, err := http.Head(url)
		if err != nil {
			fmt.Print("error ")
		}
		fmt.Printf("status =%s\n", req.Status)
	}
}
