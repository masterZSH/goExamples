package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello/", helloFunc)
	http.HandleFunc("/shuthello/", shouthelloFunc)
	http.ListenAndServe("localhost:9999", nil)
}

func helloFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello "+req.URL.Path[len("/hello/"):])
}

func shouthelloFunc(w http.ResponseWriter, req *http.Request) {
	res := req.URL.Path[len("/shuthello/"):]
	res = strings.ToUpper(res)
	fmt.Fprintf(w, "hello "+res)
}
