package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// var url = "https://v2ex.com"
	var errUrl = "https://vvvv.v2.co"
	res, err := http.Get(errUrl)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Print(string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
