package main

import "fmt"

func main() {

	defer func() {
		if a := recover(); a != nil {
			fmt.Printf("123%s", a)
		}
	}()
	i, j := 10, 0
	r := i / j
	fmt.Print(r)
}
