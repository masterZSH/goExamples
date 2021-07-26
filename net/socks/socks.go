package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	HttpProxy  = "http://127.0.0.1:10809"
	SocksProxy = "socks5://127.0.0.1:8000"
)

func main() {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(HttpProxy)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}

	req, err := http.NewRequest("GET", "https://api.ip.sb/ip", nil)
	if err != nil {
		// handle error
		panic(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	proxy = func(_ *http.Request) (*url.URL, error) {
		return url.Parse(SocksProxy)
	}

	httpTransport = &http.Transport{
		Proxy: proxy,
	}

	httpClient = &http.Client{
		Transport: httpTransport,
	}

	req, err = http.NewRequest("GET", "https://api.ip.sb/ip", nil)
	if err != nil {
		// handle error
	}

	resp, err = httpClient.Do(req)
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
