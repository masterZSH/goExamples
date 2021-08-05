package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	trace := &httptrace.ClientTrace{
		ConnectStart: func(network, addr string) {
			fmt.Println(network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println(network, addr, err)
		},
		GetConn: func(hostPort string) {
			fmt.Println(hostPort)
		},
		GotFirstResponseByte: func() {

		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Printf("%+v", info)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
