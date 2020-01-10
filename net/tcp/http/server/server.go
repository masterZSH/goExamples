package main

import "net/http"

import "fmt"

// Hello empty struct
type Hello struct{}

// 任意类型带有ServeHTTP方法即可 即存在Hello.ServeHTTP
func (h Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:8888", h)
}
