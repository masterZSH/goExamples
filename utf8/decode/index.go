package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bs := []byte{' '}
	a, i := utf8.DecodeRune(bs)
	fmt.Print(a, i)
}
