package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	{
		c := make(chan os.Signal, 1)
		// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		signal.Notify(c)

		// Block until a signal is received.
		s := <-c
		fmt.Println("Got signal:", s)
	}
}
