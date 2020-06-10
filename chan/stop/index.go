package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	urls []string = []string{
		"https://v2ex.com/",
		"https://github.com",
	}
)

func main() {
	signStop := make(chan bool, 1)
	for _, item := range urls {
		go fetch(signStop, item)
	}
	time.Sleep(3 * time.Second)
	signStop <- true
	time.Sleep(60 * time.Second)
}

func fetch(signStop chan bool, url string) {
	select {
	case <-signStop:
		fmt.Println("结束")
		return
	default:
		fmt.Printf("请求%s...\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("错误%v\n", err)
		}
		// buf := new([4096]byte)
		// _, err = resp.Body.Read(buf[:])
		// resp.Body.Close()
		fmt.Printf("读取code%d\n", resp.StatusCode)
	}
}
