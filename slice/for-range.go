package main

import (
	"fmt"
)

func main() {
	// var s = slice4([]int{1, 2, 3}...)
	// fmt.Print(s)

	// var s = slice5([]float32{1.1, 2.2, 3.3})
	// fmt.Print(s)
	// var si, af, sf, ai = slice6(1, 2.2)
	// fmt.Print(si, af, sf, ai)
	var arr = []int{1, 2, 3, 4, 5, 6}
	var sMin = slice7(arr, "min")
	var sMax = slice7(arr, "max")
	fmt.Print(sMin, sMax)

}

func slice1() {
	var s1 = make([]int, 4)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	s1[3] = 4
	for idx, v := range s1 {
		fmt.Printf("k=%d,v=%d\n", idx, v)
	}
}

func slice2() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	var season string
	// 使用_忽略索引
	for _, season = range seasons {
		fmt.Printf("Season  is: %s\n", season)
	}

	// 使用_忽略季节值
	var idx int
	for idx, _ = range seasons {
		fmt.Printf("Season index  is: %d\n", idx)
	}
}

func slice3() {
	// 修改切片的值
	var items = [...]int{10, 20, 30, 40, 50}
	for k, item := range items {
		items[k] = item * 2
	}

	for _, item := range items {
		fmt.Printf("item=%d\n", item)
	}
}

func slice4(s1 ...int) (r int) {
	// 问题7.6练习 通过使用省略号操作符 ... 来实现累加方法。
	for _, v := range s1 {
		r += v
	}
	return
}

func slice5(arrF []float32) (r float32) {
	// 不考虑溢出
	for _, v := range arrF {
		r += v
	}
	return
}

// 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。
func slice6(i int, f float32) (sumi int, averagef float32, sumf float32, averagei int) {
	sumi = i + int(f)
	sumf = float32(i) + f
	averagef = sumf / 2
	averagei = sumi / 2
	return
}

// 练习 写一个 minSlice 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice 方法返回最大值。

/**
* @param {[]int} arr 需要查找的切片
* @param {string} gType  类型 min|max 最小|最大
* @returns {int} r
 */
func slice7(arr []int, gType string) (r int) {
	// 初始化
	r = arr[0]
	switch gType {
	case "min":
		for _, v := range arr {
			if v < r {
				r = v
			}
		}
	case "max":
		for _, v := range arr {
			if v > r {
				r = v
			}
		}
	}
	return
}
