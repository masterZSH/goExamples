package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	kb
	mb
)

func main() {
	fmt.Println(kb, mb)
}
