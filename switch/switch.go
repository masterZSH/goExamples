package main

import "fmt"

func main() {
	fmt.Println(even(2))
	fmt.Println(even(21))

}

func even(x int) bool {
	switch {
	case x%2 == 0:
		return true
	case x%2 != 0:
		return false
	default:
		return false
	}
}
