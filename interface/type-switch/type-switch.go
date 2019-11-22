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

type SimpleG struct {
	i int
}

func (s *SimpleG) Get() int {
	return s.i + 1
}

func (s *SimpleG) Set(i int) {
	s.i = i + 1
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

	if t, ok := s.(*Simple); ok {
		fmt.Printf("The type of Simpler is: %T\n", t)
	}

	switch t := s.(type) {
	case *Simple:
		fmt.Printf("Type Simple %T with value %v\n", t, t)
	case *SimpleG:
		fmt.Printf("Type SimpleG %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	fmt.Printf("%+v", s)
}
