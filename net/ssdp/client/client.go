package main

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/ipv4"
)

type Discovered struct {
	// Address is the local address of a discovered peer.
	Address string
	// Payload is the associated payload from discovered peer.
	Payload []byte
}

// peer discovery configure.
type Config struct {
	// Limit is the number of peers to discover, use < 1 for unlimited.
	Limit int
	// Port is the port to broadcast on (the peers must also broadcast using the same port).
	// The default port is 9999.
	Port uint16
	// MulticastAddress specifies the multicast address.
	// You should be able to use any of 224.0.0.0/4 or ff00::/8.
	// By default it uses the Simple Service Discovery Protocol
	// address (239.255.255.250 for IPv4 or ff02::c for IPv6).
	MulticastAddress string
	// Payload is the bytes that are sent out with each broadcast. Must be short.
	Payload []byte
	// PayloadFunc is the function that will be called to dynamically generate payload
	// before every broadcast. If this pointer is nil `Payload` field will be broadcasted instead.
	PayloadFunc func() []byte
	// Delay is the amount of time between broadcasts. The default delay is 1 second.
	Delay time.Duration
	// TimeLimit is the amount of time to spend discovering, if the limit is not reached.
	// A negative limit indiciates scanning until the limit was reached or, if an
	// unlimited scanning was requested, no timeout.
	// The default time limit is 10 seconds.
	TimeLimit time.Duration
	// StopChan is a channel to stop the peer discvoery immediatley after reception.
	StopChan chan interface{}
	// AllowSelf will allow discovery the local machine (default false)
	AllowSelf bool
	// DisableBroadcast will not allow sending out a broadcast
	DisableBroadcast bool
	// Notify will be called each time a new peer was discovered.
	// The default is nil, which means no notification whatsoever.
	Notify func(Discovered)

	portNum                 int
	multicastAddressNumbers net.IP
}

type Discovery interface{}

type IPV4Discovery struct{}
type IPV6Discovery struct{}

func NewIPV4Discovery() *IPV4Discovery {
	d := &IPV4Discovery{}

	return d
}

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
