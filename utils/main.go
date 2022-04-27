package main

import (
	"fmt"

	funk "github.com/thoas/go-funk"
)

func main() {
	// Contains 包含
	t1 := funk.Contains([]string{"foo", "bar"}, "bar") // true
	fmt.Println(t1)
	// Intersect 交集
	t2 := funk.Intersect([]int{1, 2, 3, 4}, []int{2, 4, 6}) // []int{2, 4}
	fmt.Println(t2)

	// Difference 差异
	t3, t4 := funk.Difference([]int{1, 2, 3, 4}, []int{2, 4, 6})
	fmt.Println(t3, t4)

	// IndexOf 位置
	t5 := funk.IndexOf([]string{"foo", "bar"}, "bar") // 1
	fmt.Println(t5)

	// Keys 键
	t6 := funk.Keys(map[string]int{"one": 1, "two": 2}) // []string{"one", "two"} (iteration order is not guaranteed)
	fmt.Println(t6)

	// Unip  去重
	t7 := funk.Uniq([]int{0, 1, 1, 2, 3, 0, 0, 12}) // []int{0, 1, 2, 3, 12}
	fmt.Println(t7)

	// Subtract 差
	t8 := funk.Subtract([]int{0, 1, 2, 3, 4}, []int{0, 4}) // []int{1, 2, 3}
	fmt.Println(t8)

}
