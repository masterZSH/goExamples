package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 测试内存占用 堆分配内存测试

	// 空结构体
	// heap alloc 2.8MB
	mp := make(map[string]struct{})
	if _, ok := mp["test"]; !ok {
		mp["test"] = struct{}{}
	}

	// heap alloc 3.2MB
	// mp := make(map[string]bool)
	// if _, ok := mp["test"]; !ok {
	// 	mp["test"] = false
	// }

	log.Println(http.ListenAndServe("localhost:6060", nil))

}
