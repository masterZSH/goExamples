package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", helloFunc)
	http.ListenAndServe("localhost:9999", nil)
}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello "+req.URL.Path[len("/hello/"):])
}
