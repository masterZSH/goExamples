package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string = "test"
	st = strings.Map(up, st)
	fmt.Printf("%s\n", st)
}

func up(r rune) rune {
	return r + 'A' - 'a'
}
