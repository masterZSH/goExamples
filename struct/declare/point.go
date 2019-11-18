package main

import (
	"fmt"
	"math"
)

// Point 2d
type Point struct {
	x float64
	y float64
}

// Polar 3d
type Polar struct {
	z int
	p *Point
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func Scale(p *Point, s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func main() {
	var p = new(Point)
	p.x = 1.1
	p.y = 2.2
	fmt.Printf("%+v", p)
	fmt.Printf("%f", Abs(p))
	Scale(p, 2.2)
	fmt.Printf("%+v", p)
}
