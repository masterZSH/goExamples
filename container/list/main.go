package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(1)
	l.PushFront(2)
	fmt.Println(l.Len())
}
