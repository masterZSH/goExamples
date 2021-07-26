package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func main() {
	// body := new(bytes.Buffer)
	// mw := multipart.NewWriter(body)
	// defer mw.Close()
	// mw.WriteField("foo", "bar")
	// boundary := "--testboundary"
	// mw.SetBoundary(boundary)

	// req, _ := http.NewRequest("POST", "/", body)
	// req.Header.Set("Content-Type", "multipart/form-data"+"; boundary="+boundary)

	// err := req.ParseMultipartForm(33554432)
	// fmt.Println(err)
	// err = req.ParseForm()
	// fmt.Println(err)
	req := newRequest()
	fmt.Println(req)
	err := req.ParseMultipartForm(33554432)
	fmt.Println(err)
	fmt.Println(req.MultipartForm)
}

func newRequest() *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	mw.SetBoundary(boundary)
	mw.WriteField("foo", "bar")
	mw.WriteField("bar", "10")
	mw.WriteField("bar", "foo2")
	req, err := http.NewRequest("POST", "/", body)
	fmt.Println(err)
	req.Header.Set("Content-Type", "multipart/form-data"+"; boundary="+boundary)
	return req
}
