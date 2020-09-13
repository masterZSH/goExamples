package main

import (
	"fmt"
	"sync"
)

// Map is like a Go map[interface{}]interface{}
func main() {
	m := new(sync.Map)
	m.Store(1, 2)
	fmt.Println(m.Load(1))
}
