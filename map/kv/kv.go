package main

import (
	"fmt"
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
}
