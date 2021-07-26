package main

import (
	"context"
	"fmt"
	"net"
)

func main() {
	rs := net.Resolver{
		PreferGo: true,
		Dial:     dial,
	}
	// ip, err := rs.LookupIP(context.Background(), "ip", "myip.opendns.com")
	ip, err := rs.LookupIP(context.Background(), "ip", "www.google.com")
	fmt.Println(ip, err)
}

var dialNetworkSuffixVar = &struct{}{}

func dial(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{
		// I think we want to disable switching to other networks, since this DNS response
		// depends on the network used.
		FallbackDelay: -1,
	}
	// Go DNS resolution doesn't bother to use a particular network, but we need to reapply our
	// constraint.
	suffix := ctx.Value(dialNetworkSuffixVar)
	if suffix != nil {
		network += suffix.(string)
	}
	return d.DialContext(ctx, network, "resolver1.opendns.com:53")
}
