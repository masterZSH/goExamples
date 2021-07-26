package main

import (
	"fmt"
	"net"
)

func main() {
	ipv4Addr := net.ParseIP("192.0.2.1")
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))

	// ipv6 mask
	ipv6Addr := net.ParseIP("fe80::14a1:40d4:952c:bd16")
	ipv6Mask := net.CIDRMask(64, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))

}
