package main

import (
	"fmt"
)

type day int

const (
	Mo day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%f/%s", t.a, t.b, t.c)
}

func main() {
	// var t *T = &T{1, 1.2, "12"}
	// fmt.Print(t)
	var c Celsius = 1.2
	fmt.Print(c)
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%d%s", int(c), "°C")
}
