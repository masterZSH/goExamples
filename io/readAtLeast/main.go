package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	rd := strings.NewReader("abcde")
	buf := make([]byte, 4)

	n, err := io.ReadAtLeast(rd, buf, 2)
	fmt.Printf("%d,%v,%s\n", n, err, buf[:n])

	// rd length > min
	rd = strings.NewReader("abcde")
	buf = make([]byte, 10)
	// rd len
	n, err = io.ReadAtLeast(rd, buf, 6)
	// 5,unexpected EOF,abcde
	fmt.Printf("%d,%v,%s\n", n, err, buf[:n])

	// min > len(buf)
	rd = strings.NewReader("abcde")
	buf = make([]byte, 4)
	n, err = io.ReadAtLeast(rd, buf, 6)
	// 0,short buffer,
	fmt.Printf("%d,%v,%s\n", n, err, buf[:n])

}
