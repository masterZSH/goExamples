
package main

import "fmt"

func main() {
	m := map[string]int64{
		"1": 1,
		"2": 2,
	}
	s := SumIntsOrFloats(m)
	// 详细写法
	// s := SumIntsOrFloats[string, int64](m)

	fmt.Println(s)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
