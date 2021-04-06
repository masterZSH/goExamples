package main

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	//
	fmt.Println(ifaces)
	c, err := net.ListenPacket("udp4", "239.255.255.250:6666")
	fmt.Println(c, err)

	cn := ipv4.NewPacketConn(c)
	go listen(cn)

	ip := net.ParseIP("239.255.255.250")
	addr := &net.UDPAddr{IP: ip, Port: 6666}
	for i := range ifaces {
		cn.JoinGroup(&ifaces[i], addr)
	}

	n, err := cn.WriteTo([]byte("test"), nil, addr)
	fmt.Println(n, err)
	for {

	}
}

func listen(con *ipv4.PacketConn) {
	for {
		buffer := make([]byte, 1024)
		n, _, src, errRead := con.ReadFrom(buffer)
		fmt.Println(n, src, errRead)
	}
}
