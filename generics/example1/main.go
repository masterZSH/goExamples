package main

import "fmt"

type Foo struct {
	Name string
}

type Bar struct {
	Name string
}

func main() {
	arr := []int{1, 2, 3, 4}
	tArr := Del(arr, 3)
	fmt.Println(tArr)

	fooArr := []Foo{
		Foo{"1"},
		Foo{"2"},
		Foo{"3"},
		Foo{"4"},
	}

	fooArr = Del(fooArr, Foo{"3"})
	fmt.Println(fooArr)

	// pointer
	b1 := &Bar{"1"}
	b2 := &Bar{"2"}
	b3 := &Bar{"3"}
	b4 := &Bar{"4"}
	barArr := []*Bar{
		b1,
		b2,
		b3,
		b4,
	}
	fmt.Println(barArr)

	barArr = Del(barArr, b3)
	fmt.Println(barArr)

}

func Del[T comparable](src []T, v T) []T {
	if len(src) == 0 {
		return src
	}
	delIdx := -1
	for i := range src {
		if src[i] == v {
			delIdx = i
			break
		}
	}
	if delIdx == -1 {
		return src
	}
	return append(src[:delIdx], src[delIdx+1:]...)
}
