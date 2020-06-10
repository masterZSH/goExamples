package main

import "fmt"

func main() {
	db(5)
}

//
func db(i int) (result int) {
	defer func() {
		fmt.Printf("%d\n", i)
	}()
	return i + i
}
