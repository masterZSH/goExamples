package main

import "fmt"

func main() {
	m1 := map[string]int64{
		"1": 1,
		"2": 2,
	}
	s1 := SumIntsOrFloats(m1)
	// 详细写法
	// s := SumIntsOrFloats[string, int64](m)

	fmt.Println(s1)

	m2 := map[string]float64{
		"1": 1.1,
		"2": 2.2,
	}
	s2 := SumIntsOrFloats(m2)
	fmt.Println(s2)
}

// comparable is an interface that is implemented by all comparable types
// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint,
// not as the type of a variable.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
