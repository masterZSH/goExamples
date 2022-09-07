package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	f1()
}

func f1() {
	f2()
}

func f2() {
	fmt.Printf("%s\n", debug.Stack())
}
