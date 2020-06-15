package main

import (
	"fmt"
	"sync"
)

var (
	i, j int
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		i = 1
		fmt.Printf("j = %d\n", j)
		wg.Done()
	}()
	go func() {
		j = 1
		fmt.Printf("i = %d\n", i)
		wg.Done()
	}()
	wg.Wait()

}
