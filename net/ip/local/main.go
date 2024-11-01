package main

import (
	"fmt"
	"net"
)

var v4Loopback = net.IPNet{
	IP:   net.IP{127, 0, 0, 0},
	Mask: net.IPv4Mask(255, 0, 0, 0),
}

func ipIsLocal(ip net.IP) bool {
	if ip.To4() == nil {
		return ip.Equal(net.IPv6loopback)
	}

	return v4Loopback.Contains(ip)
}

func main() {
	ip := net.ParseIP("192.168.1.21")
	r := v4Loopback.Contains(ip)
	fmt.Println(r)
}
