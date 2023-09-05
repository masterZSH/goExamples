package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	testStr := "123世界"
	for len(testStr) > 0 {
		r, size := utf8.DecodeRuneInString(testStr)
		fmt.Printf("%c %v\n", r, size)

		testStr = testStr[size:]
	}
	// output
	// 1 1
	// 2 1
	// 3 1
	// 世 3
	// 界 3

}
