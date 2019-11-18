package main

import (
	"fmt"
)

type Rectangle struct {
	width, height int
}

// Area 面积
func Area(r *Rectangle) int {
	return r.width * r.height
}

// Perimeter 周长
func Perimeter(r *Rectangle) int {
	return 2 * (r.width + r.height)
}

func main() {
	// test Area
	var r Rectangle
	r.width = 2
	r.height = 3
	fmt.Printf("%d\n", Area(&r))
	fmt.Printf("%d\n", Perimeter(&r))

}
