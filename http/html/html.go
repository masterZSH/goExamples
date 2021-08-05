package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	res, err := http.Get("http://baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	root, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root.FirstChild.Data)
}
