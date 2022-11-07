package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, 1, 1)
	fmt.Println(fd, err)
}
