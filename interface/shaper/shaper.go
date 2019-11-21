package main

import "fmt"

type Square struct {
	side float64
}

type Shaper interface {
	Area() float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var r = new(Rectangle)
	r.width = 5
	r.height = 4
	var s = &Square{3}
	shapers := []Shaper{r, s}
	for _, v := range shapers {
		fmt.Println(v.Area())
	}

}
