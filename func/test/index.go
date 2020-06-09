package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	arr := []int{0, 1, 2}
	fmt.Println(test(arr, 3))
	fmt.Println(test(arr, 4))
	http.ListenAndServe(":6060", nil)
}

func test(arr []int, s ...int) []int {
	return append(arr, s...)
}
