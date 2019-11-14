package main

import (
	"fmt"
	"sync"
)

type Info struct {
	mu sync.Mutex
	// name
	name string
}

func main() {
	var i = Info{}
	update(&i)
	fmt.Print(i.name)
}

// 这是线程安全的 因为借助了互斥锁
func update(i *Info) {
	i.mu.Lock()
	i.name = "123"
	i.mu.Unlock()
}

// RWMutex锁读写锁
