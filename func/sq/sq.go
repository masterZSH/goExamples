package main

import "fmt"

func main() {
	f := sq()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func sq() func() int {
	var i int
	return func() int {
		i++
		return i * i
	}
}
