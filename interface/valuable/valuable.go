package main

import (
	"fmt"
)

type Stock struct {
	s int
}

func (s *Stock) getValue() int {
	return s.s
}

type Car struct {
	price int
	other int
}

func (c *Car) getValue() int {
	return c.price + c.other
}

type valuable interface {
	getValue() int
}

func showValue(v valuable) {
	fmt.Println(v.getValue())
}

func main() {
	var v1 valuable = &Car{1, 2}
	var v2 valuable = &Stock{10}
	showValue(v1)
	showValue(v2)
}
