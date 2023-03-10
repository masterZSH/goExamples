package main

import "fmt"

type Tm string

func main() {
	s1 := "1"
	s2 := Tm("1")

	St(s1)
	St(s2)
}

func St[T ~string](t T) {
	fmt.Println(t)
}
