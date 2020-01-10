package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", testFunc)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("error start server")
	}
}

func testFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
