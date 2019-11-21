package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(i int) {
	s.i = i
}

func main() {
	var s Simpler = &Simple{12}
	fmt.Println(s.Get())
	s.Set(122)
	fmt.Println(s.Get())

	fmt.Printf("%+v", s)
}
