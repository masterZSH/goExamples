package main

import (
	"errors"
	"fmt"
	"strings"
	"testing/iotest"
)

func main() {
	// A reader that always returns a custom error.
	r := iotest.ErrReader(errors.New("custom error"))
	n, err := r.Read(nil)
	fmt.Printf("n:   %d\nerr: %q\n", n, err)

	// half reader
	rd := strings.NewReader("foo bar")
	halfRd := iotest.HalfReader(rd)
	bf := make([]byte, 6)
	halfRd.Read(bf)
	fmt.Printf("%s\n", bf)
}
