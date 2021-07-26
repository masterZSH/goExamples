package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func s() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")

	})
	serv := httptest.NewServer(handler)
	defer serv.Close()
	client := serv.Client()
	resp, err := client.Get(serv.URL)
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%q", data)

}
