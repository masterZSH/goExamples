package main

import "fmt"

func main() {
	var map1 map[string]int = make(map[string]int)
	map1["1"] = 1
	fmt.Printf("%+v\n", map1)
	// for range

	for k, v := range map1 {
		fmt.Print(k, v)
	}

	if val, isPresent := map1["1"]; isPresent {
		fmt.Print(val)
	}
	// 不存在

	// 删除
	map1["222"] = 3
	delete(map1, "1")
	fmt.Printf("%+v", map1)
}
