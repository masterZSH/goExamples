package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Unix())
	fmt.Println(t.Format(time.Stamp))

}
