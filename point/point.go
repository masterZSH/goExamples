package main

import "fmt"

// *Point(p)        // same as *(Point(p))
// (*Point)(p)      // p is converted to *Point
// <-chan int(c)    // same as <-(chan int(c))
// (<-chan int)(c)  // c is converted to <-chan int
// func()(x)        // function signature func() x
// (func())(x)      // x is converted to func()
// (func() int)(x)  // x is converted to func() int
// func() int(x)    // x is converted to func() int (unambiguous)

func main() {
	i := 1
	f := *float64(&i)
	fmt.Print(&f)

}
