package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", test1Handler)
	r.HandleFunc("/2", test2Handler)
	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}

func test1Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test1"))
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test2"))

}
