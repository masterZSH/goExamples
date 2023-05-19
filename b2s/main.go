package main

import (
	"fmt"
	"unsafe"
)

// go 1.20
// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
func b2s(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func main() {
	str := b2s([]byte{97, 98, 99})
	fmt.Println(str)
}
