package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "/?foo=bar", nil)
	qs := req.URL.Query()
	fmt.Println(qs)
}
