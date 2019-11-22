package main

type A interface {
	print()
}

type B interface{}

func p(a A) {

}

func main() {
	var b B
	b.(A).print()
}
