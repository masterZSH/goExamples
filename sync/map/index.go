package main

import (
	"fmt"
	"sync"
)

// Map is like a Go map[interface{}]interface{}
func main() {
	var m sync.Map
	// m1 := make(map[int]string)

	go func() {
		for {
			m.Store(1, "val")
			// m1[1] = "val"
		}
	}()
	go func() {
		// 不停地对map进行读取
		for {
			// _ = m1[1]
			fmt.Println(m.Load(1))
		}
	}()

	for {

	}

}
