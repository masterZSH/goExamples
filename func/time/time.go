package main

import (
	"fmt"
	"time"
)

type mTime struct {
	t time.Time
}

// error
// func (t *time.Time)pTime(){
// 	fmt.Print(t.Date())
// }

func (t *mTime) pTime() {
	fmt.Print(t.t.Date())
}

func main() {
	var t = new(mTime)
	t.t = time.Now()
	t.pTime()
}
