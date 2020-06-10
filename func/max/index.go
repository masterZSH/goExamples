package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("%d\n", m(1, 2, 3, 4, 56))

	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic: %v", e)
		}
	}()
	fmt.Printf("%d\n", m())
}

func m(arg ...int) int {
	if len(arg) == 0 {
		panic("no args")
	}
	var a int
	for _, v := range arg {
		if v > a {
			a = v
		}
	}
	return a
}
