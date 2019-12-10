package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main")
	go ShortWait()
	go LongWait()
	fmt.Println("start ")
	time.Sleep(10 * 1e9)
	fmt.Println("end ")
}

// ShortWait sleep 2 seconds
func ShortWait() {
	fmt.Println("In ShortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("ShortWait end")
}

// LongWait sleep 5 seconds
func LongWait() {
	fmt.Println("In LongWait")
	time.Sleep(5 * 1e9)
	fmt.Println("LongWait end")
}
