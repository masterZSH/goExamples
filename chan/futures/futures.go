package main

import "fmt"

func Product(a, b int) int {
	da_future := Double(a)
	db_future := Double(b)
	da := <-da_future
	db := <-db_future
	return da + db
}

func Double(a int) chan int {
	ch := make(chan int)
	go func() {
		ch <- a * 2
	}()
	return ch
}

func main() {
	fmt.Print(Product(123, 234))
}
