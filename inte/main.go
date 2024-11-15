package main

type I interface {
	f1()
	f2()
}

type T struct {
}

func (t *T) f1() {
	println("f1")
}

func (t *T) f2() {
	println("f2")
}

var _ I = (*T)(nil)

func main() {
	t := new(T)
	t.f1()
	t.f2()
}
