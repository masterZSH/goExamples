package main

import (
	"fmt"

	snappy "github.com/klauspost/compress/s2"
)

func main() {
	// compress

	src := []byte("sxccvb123234--12")
	result := snappy.Encode(nil, src)
	fmt.Printf("%s\n", result)

	// decode
	result, _ = snappy.Decode(nil, result)
	fmt.Printf("%s\n", result)
}
