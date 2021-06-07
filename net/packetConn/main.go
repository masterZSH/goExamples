package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

func main() {

	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	pc, err := net.ListenPacket("udp4", "239.255.255.250:9362")
	if err != nil {
		return
	}

	packetConn := ipv4.NewPacketConn(pc)

	for _, ifi := range ifaces {
		packetConn.JoinGroup(
			&ifi,
			&net.UDPAddr{IP: net.ParseIP("239.255.255.250"), Port: 9362},
		)
	}
	fmt.Printf("%q", ifaces)

	go func() {
		time.Sleep(5 * time.Second)
		packetConn.WriteTo(
			[]byte("abc"),
			nil,
			&net.UDPAddr{IP: net.ParseIP("239.255.255.250"), Port: 9362},
		)
	}()

	for {
		bs := make([]byte, 1024)
		n, msg, src, err := packetConn.ReadFrom(bs)
		if err != nil {
			fmt.Printf("err %s\n", err)
		}
		fmt.Printf("%d,content:%s,msg:%q,src:%s\n", n, bs[:n], msg, src)
	}

}
