package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/cloudwego/hertz/pkg/protocol/http1/resp"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("listening on %s\n", l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	rd := bufio.NewReader(c)
	rq, err := http.ReadRequest(rd)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", rq)
	if rq.Method == http.MethodConnect {
		// Manual writing to support CONNECT for http 1.0 (workaround for uplay client)
		if _, err = fmt.Fprintf(c, "HTTP/%d.%d %03d %s\r\n\r\n", rq.ProtoMajor, rq.ProtoMinor, http.StatusOK, "Connection established"); err != nil {
			return // close connection
		}
	}

	rc, err := net.Dial("tcp", rq.Host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()
	rc.Write(c)
	go io.Copy(remoteConn, c.reqConn)

	io.Copy(c.reqConn, remoteConn)

	// rq.RequestURI = ""
	// rq.URL.Scheme = "http"
	// rq.URL, _ = url.Parse(rq.URL.String())
	// rq.URL.Host = rq.Host
	// resp, err := http.DefaultClient.Do(rq)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	resp.Write(c)
}
