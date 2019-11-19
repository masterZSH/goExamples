package main

import (
	"fmt"
	"strconv"
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
	var t *T = &T{1, 1.2, "12"}
	fmt.Print(t)
}

type Celsius float64

func (c *Celsius) String() string  {
	
}
