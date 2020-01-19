package main

import "fmt"

import "time"

func main() {
	var arr []int = []int{10, 20, 30, 40}
	for ix := range arr {
		func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()

	for ix := range arr {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)

	for ix := range arr {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)
}
