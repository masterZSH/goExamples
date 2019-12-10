package main

import "fmt"
import "time"

func main() {
	// chanString := make(chan string)
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "string1"
	ch <- "string2"
	ch <- "string3"
	ch <- "string4"
}

func getData(ch chan string) {
	var i string
	for {
		i = <-ch
		fmt.Printf("%s", i)
	}
}
