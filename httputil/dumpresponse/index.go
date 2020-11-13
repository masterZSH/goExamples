package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

const body = "test"

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		fmt.Printf("%+v,%+v\n", w, q)
		w.Write([]byte(body))
	}))
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)
}
