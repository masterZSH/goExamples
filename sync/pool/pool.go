package main

import (
	"fmt"
	"sync"
)

var (
	pool = sync.Pool{
		New: func() interface{} {
			return make(map[string]int)
		},
	}
)

func main() {
	go t()
	defer func() {
		t()
	}()
}

func t() {
	item := pool.Get().(map[string]int)
	fmt.Printf("%v", item)
	pool.Put(item)
}
