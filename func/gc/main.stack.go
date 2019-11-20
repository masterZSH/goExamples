package main

import (
	"stack"
	"fmt"
)



func main() {
	var s = new(stack.Stack)
	s.Push(1)
	fmt.Println(s)
}
