package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	fmt.Println(err)
	for _, addr := range addrs {
		fmt.Println(addr)
	}

}
