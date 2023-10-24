package main

import "time"

func main() {
	// dead lock
	tm := time.NewTimer(1)
	// tm.Reset(100 * time.Millisecond)
	go func() {
		time.Sleep(5 * time.Second)
		tm.Reset(100 * time.Millisecond)

	}()
	<-tm.C
	if !tm.Stop() {
		<-tm.C
	}
}
