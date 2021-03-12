package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	mp      sync.Map
	clients sync.Map
)

func init() {
	mp = sync.Map{}
	clients = sync.Map{}
}

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	log.Printf("conn %s\n", conn.RemoteAddr())
	mp.Store(conn.RemoteAddr(), 1)
	go heartBeat(conn)
	buf := make([]byte, 1024)
	for {
		alive, ok := mp.Load(conn.RemoteAddr())
		if !ok {
			log.Printf("load conn err\n")
			return
		}
		if alive.(int) == 0 {
			log.Printf("load conn err\n")
			return
		}
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("get message error:%s", err)
			return
		}
		log.Printf("get msg %s\n", buf[:n])
		v, _ := mp.Load(conn.RemoteAddr())
		log.Printf("conn:%+v\n", v)
		if bytes.Equal(buf[:n], []byte("ping")) {
			conn.Write([]byte("pong"))
			mp.Store(conn.RemoteAddr(), 1)
			continue
		}
		// get-name
		if n < 4 {
			continue
		}

		if bytes.Equal(buf[:3], []byte("get")) {
			if v, ok := clients.Load(string(buf[3:n])); ok {
				conn.Write([]byte(v.(string)))
			}
			continue
		}
		// store client
		log.Printf("%s store %s\n", buf[:n], conn.RemoteAddr())
		clients.Store(string(buf[:n]), conn.RemoteAddr().String())
	}

}

func heartBeat(conn net.Conn) {
	tk := time.NewTicker(time.Minute)
	defer tk.Stop()
	if _, ok := mp.Load(conn.RemoteAddr()); ok {
		for {
			select {
			case _ = <-tk.C:
				fmt.Println("tk")
				mp.Store(conn.RemoteAddr(), 0)
				clients.Range(func(k, v interface{}) bool {
					if v.(string) == conn.RemoteAddr().String() {
						clients.Delete(k)
					}
					return true
				})

			}
		}
	}
}

func ping() {

}

func handshake() {

}
