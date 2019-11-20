package main

import (
	"fmt"
)

const LIMIT = 11

// Stack 后进先出
type Stack struct {
	i    int
	data [LIMIT]int
}

func (s *Stack) Push(n int) {
	if s.i+1 > LIMIT {
		return // stack is full!
	}
	s.data[s.i] = n
	s.i++
}

func (s *Stack) Pop() int {
	defer func() {
		s.data[s.i-1] = 0
		s.i--
	}()
	return s.data[s.i-1]
}

func (s *Stack) String() {
	fmt.Printf("%+v", s)
}

func main() {
	var s = new(Stack)
	s.Push(1)
	fmt.Println(s)

	s.Push(2)
	fmt.Println(s)

	s.Push(3)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)

}
