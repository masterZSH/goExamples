package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int, 1)
	map1[0] = 1
	map1[1] = 2
	// v-值  present-key1是否在map中
	v, present := map1[1]
	fmt.Print(v, present)
	// 删除key
	fmt.Print(map1)
	delete(map1, 1)
	fmt.Print(map1)
	// 配合if书写
	if v, present := map1[0]; present {
		fmt.Print(v)
	}

	//
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	// range
	for k, v := range capitals {
		fmt.Printf("k=%s;v=%s\n", k, v)
	}

	week := map[int]string{1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday"}

	for k, v := range week {
		fmt.Printf("k=%d;v=%s\n", k, v)
	}

	var ma = map[string]int{"foo": 1, "bar": 2, "z": 3, "abc": 4}
	sortMap(ma)

	drink()
}

type name struct {
	key   string
	value int
}

type names []name

func sortMap(m map[string]int) {
	var n names = []name{}
	for k, v := range m {
		n = append(n, name{
			key:   k,
			value: v,
		})
	}
	// sort.Strings(n)
	fmt.Print(n)
}

// 构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
// 先打印所有的饮料，然后打印原名和翻译后的名字。
// 接下来按照英文名排序后再打印出来。
func drink() {
	var m = map[string]string{
		"cola":  "可乐",
		"water": "水",
		"tea":   "茶",
	}
	for k, v := range m {
		fmt.Printf("drank is %s\n", k)
		fmt.Printf("%s,%s\n", k, v)
	}
	// sort
	var s = []string{}
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, v := range s {
		fmt.Printf("%s,%s\n", v, m[v])
	}

}
