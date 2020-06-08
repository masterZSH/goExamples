package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["t1"] = "123"
	mp["t2"] = "456"
	// map[t1:123 t2:456]
	fmt.Printf("%+v\n",mp) 
	delete(mp,"t1")
	// map[t2:456]
	fmt.Printf("%+v\n",mp) 

	// 
	if v,ok := mp["t3"]; !ok{
		fmt.Print(v)
	}
}