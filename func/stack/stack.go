package main

import (
	"fmt"
)

// Stack 后进先出
type Stack [4]int

func (s *Stack) Push(n int) {
	for k, v := range s {
		if v == 0 {
			s[k] = n
			break
		}
	}
}

func (s *Stack) Pop() int {
	for ix := len(s) - 1; ix >= 0; ix-- {
		if s[ix] != 0 {
			defer func() {
				s[ix] = 0
			}()
			return s[ix]
		}
	}
	return 0
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
