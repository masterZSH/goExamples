package main

import (
	"fmt"
)

func main() {
	var sl = make([]byte, 1)
	sl = Append(sl, []byte{'a', 'b'})
	fmt.Println(sl)
	sl = Append(sl, []byte{'a', 'b'})
	fmt.Println(sl)

}

func Append(slice, data []byte) []byte {
	var sliceC = slice
	for _, v := range data {
		sliceC = append(slice, v)
	}
	return sliceC
}
