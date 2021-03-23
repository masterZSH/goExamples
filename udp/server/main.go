package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.ListenPacket("udp", ":6668")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		buff := make([]byte, 1024)
		localAddr := conn.LocalAddr()
		fmt.Println(getLocalIPs())
		fmt.Println(localAddr)
		n, addr, err := conn.ReadFrom(buff)
		fmt.Println(buff[:n], addr, err)
	}
}

func getLocalIPs() (ips map[string]struct{}) {
	ips = make(map[string]struct{})
	ips["localhost"] = struct{}{}
	ips["127.0.0.1"] = struct{}{}
	ips["::1"] = struct{}{}

	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, address := range addrs {
			ip, _, err := net.ParseCIDR(address.String())
			if err != nil {
				// log.Printf("Failed to parse %s: %v", address.String(), err)
				continue
			}

			ips[ip.String()+"%"+iface.Name] = struct{}{}
			ips[ip.String()] = struct{}{}
		}
	}
	return
}
