package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type urlFetch struct {
	urls []string
}

var wg sync.WaitGroup

func main() {
	var u urlFetch
	u.urls = []string{
		"https://v2ex.com/",
		"https://github.com",
	}
	wg.Add(len(u.urls))
	for _, url := range u.urls {
		go fetch(url)
	}
	wg.Wait()
	fmt.Println("all goroutines done")
}

func fetch(url string) {
	resp, _ := http.Get(url)
	bosy, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", bosy)
	resp.Body.Close()
	wg.Done()
}
