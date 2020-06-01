package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s:=strings.NewReader("1233")
	var r *bufio.Reader
	r = bufio.NewReader(s)
	// 123
	st,_:= r.Peek(3)
	fmt.Printf("%s",st)
}