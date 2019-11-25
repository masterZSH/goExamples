package main

import (
	"errors"
	"fmt"
)

// MAX Stack maxlen
var MAX = 10

// Stack 练习现在用一个元素类型是 interface{}（空接口）的切片开发一个通用的栈类型。
type Stack struct {
	data []interface{}
}

// Len return Stack length
func (s *Stack) Len() int {
	return len(s.data)
}

// IsEmpty return IsEmpty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("no data")
	}
	var item = s.data[len(s.data)-1]
	defer func() {
		s.data = s.data[:len(s.data)-1]
	}()
	return item, nil
}

func main() {
	var s = &Stack{}
	s.Push(1)
	s.Push("2234")
	fmt.Printf("%+v", s)
	fmt.Print(s.Pop())

	fmt.Printf("%+v", s)

}
