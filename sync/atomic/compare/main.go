package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32
	atomic.StoreInt32(&i, 1)
	swapped := atomic.CompareAndSwapInt32(&i, 1, 1)
	fmt.Println(swapped) // true
	fmt.Println(i)       // 1

	swapped = atomic.CompareAndSwapInt32(&i, 2, 2)
	fmt.Println(swapped) // false
	fmt.Println(i)       // 1

	swapped = atomic.CompareAndSwapInt32(&i, 1, 3)
	fmt.Println(swapped) // true
	fmt.Println(i)       // 3
}
