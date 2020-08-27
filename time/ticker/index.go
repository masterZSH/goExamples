package main

import (
	"fmt"
	"time"
)

func main() {
	// duration
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			go refreshDB()
		}
	}
}

func refreshDB() {

}
