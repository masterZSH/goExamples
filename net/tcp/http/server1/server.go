package main

import "net/http"

// 简易版本

func main() {
	http.HandleFunc("/test", testFunc)
	http.ListenAndServe(":8089", nil)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("abc"))
}
