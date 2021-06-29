package main

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

func main() {
	var b1 bitset.BitSet
	b1.Set(1).Set(2)

	var b2 bitset.BitSet
	b2.Set(2).Set(3)

	count := b1.Count()
	fmt.Printf("count : %d\n", count) // 2
	b1.Set(2)
	fmt.Printf("count : %d\n", count) // 2

	// 交集
	fmt.Println(b1.Intersection(&b2))

}
