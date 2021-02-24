package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	// htpp file system 当前目录
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
