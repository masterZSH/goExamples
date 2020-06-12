package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	for i := 10; i > 0; i-- {
		fmt.Printf("%d\n", i)
		<-tick
	}

}
