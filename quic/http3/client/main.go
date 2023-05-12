package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {

	c := http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				NextProtos:         []string{"quic-test"},
			},
		},
	}
	resp, err := c.Get("https://127.0.0.1:6666")
	fmt.Println(resp, err)
}
