package main

import (
	"sync/atomic"
	"fmt"
)

func main() {
	var p atomic.Pointer[int32]
	var i int32 = 1
	p.Store(&i)
	ti:= p.Load()

	// 地址 0xc00001a0b8
	fmt.Println(ti)
	// 值  1
	fmt.Println(*ti)

	// swap 
	var ni int32 = 2
	p.Swap(&ni)
	ti= p.Load()

	// 地址 0xc00001a0f0
	fmt.Println(ti)
	// 值  2
	fmt.Println(*ti)
	
}
