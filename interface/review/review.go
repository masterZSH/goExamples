package main

import "fmt"

// S test struct
type S struct {
	name string
}

type m interface {
	string() string
}

func (s *S) string() string {
	return s.name
}

func main() {
	// 如何检测一个值v是否实现了接口Stringer：
	var v m = &S{"123"}
	if v, ok := v.(m); ok {
		fmt.Printf("%s", v.string())
	}
}
