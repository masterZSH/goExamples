package main

import "fmt"

type A interface {
	print()
}

type B interface{}

func p(a A) {

}

type M struct {
	t string
}

func (m *M) print() {
	fmt.Println("call print")
	fmt.Print(m.t)
}

func main() {
	var m B
	m = &M{"123"}
	m.(A).print()
}
