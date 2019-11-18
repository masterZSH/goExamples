package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) len() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type employee struct {
	string
	salary float64
}

// 定义结构体 employee，它有一个 salary 字段，
// 给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水。
func (e *employee) giveRaise(persent float64) {
	e.salary = e.salary * (1 + persent)
}

func main() {
	var p = &Point{3, 4}
	fmt.Printf("%f\n", p.len())
	var e = &employee{"foo", 100.1}
	e.giveRaise(0.1)
	fmt.Printf("%+v\n", e)
}
