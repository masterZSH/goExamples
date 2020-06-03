package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"sync"
)

var count int = 0

// 互斥锁
var mu sync.Mutex

func main() {
	http.HandleFunc("/", ct)
	http.HandleFunc("/count", displayCount)
	http.ListenAndServe(":8089", nil)
	fmt.Print()

}

func ct(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
}

func displayCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	f, err := os.OpenFile("1.txt", os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Fprintf(f, "count     is %d", count)
	f.Close()
	// io.Copy(w)
	f,_ = os.Open("1.txt")
	io.Copy(w,f)
	
	// fmt.Fprintf(w, "count is %d", count)
	mu.Unlock()
}

// apache benchmark 测试  -n 请求数量 -c 并发数
// ab -n 10000 -c 100 http://localhost:8089/
