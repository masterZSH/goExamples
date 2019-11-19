package main

//Integer alias int
type Integer int

func (p Integer) get() int {
	return int(p)
}

func f(i int) {

}

var v Integer

type Integer1 struct {
	n int
}

func (p *Integer1) get() int {
	return p.n
}

func f1(i int) {

}

func main() {
	f(int(v))
}
