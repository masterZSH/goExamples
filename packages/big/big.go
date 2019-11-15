package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(math.MaxInt64)
	fmt.Println(im)
	io := big.NewInt(1888)
	fmt.Printf("%v\n", io)
	fmt.Print(io.Mul(io, im))
}
