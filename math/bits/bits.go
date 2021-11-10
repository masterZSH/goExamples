package main

import (
	"fmt"
	"math/bits"
	"runtime"
)

func main() {
	fmt.Printf("Len32(%032b) = %d\n", 17, bits.Len32(17))

	r := bits.Len32(uint32(runtime.NumCPU()))

	// 8  => 1000   4 = 2**3 3+1
	fmt.Println(runtime.NumCPU(), r)
}
