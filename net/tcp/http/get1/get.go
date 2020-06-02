package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	// err := Get1(url)
	code, err := GetCode(url)
	// body, err := Get(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", code)
}

func Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return bs, err
}

func Get1(url string) error {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// 将Reader拷贝到标准输出
	ct, err := io.Copy(os.Stdout, resp.Body)
	fmt.Println(ct)
	return err
}

func GetCode(url string) (code int, err error) {
	url = getUrl(url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func getUrl(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}
