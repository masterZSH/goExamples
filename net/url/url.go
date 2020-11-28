package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://test.com/x/y")
	if err != nil {
		log.Fatal(err)
	}

	// http
	fmt.Printf("Scheme: %s\n", u.Scheme)

	// /x/y
	fmt.Println("Path:", u.Path)

	// test.com
	fmt.Println("Host:", u.Host)
}
