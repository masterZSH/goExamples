package main

import (
	"fmt"
	"time"
)

func main() {
	var x *int64 = new(int64)
	for i := 0; i < 10; i++ {
		go func(x *int64) {
			*x++
			fmt.Println(*x)
		}(x)
	}
	time.Sleep(10 * time.Second)
}
