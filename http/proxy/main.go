package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":33080")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server start %s", l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	rd := bufio.NewReader(conn)
	rq, err := http.ReadRequest(rd)
	if err != nil {
		return
	}
	rq.URL, _ = url.Parse(rq.URL.String())
	if rq.URL.Host == "" {
		rq.URL.Host = rq.Host

	}
	h, p := SplitHost(rq.URL.Host)
	host := fmt.Sprintf("%s:%s", h, p)
	rc, err := net.DialTimeout("tcp", host, 3*time.Second)
	if err != nil {
		return
	}
	if rq.Method == http.MethodConnect {
		// Manual writing to support CONNECT for http 1.0 (workaround for uplay client)
		if _, err = fmt.Fprintf(conn, "HTTP/%d.%d %03d %s\r\n\r\n", rq.ProtoMajor, rq.ProtoMinor, http.StatusOK, "Connection established"); err != nil {
			return // close connection
		}
	} else {
		rq.Write(rc)
	}

	go pipe(rc, conn)
	go pipe(conn, rc)

}

func SplitHost(host string) (string, string) {
	result, port, err := net.SplitHostPort(host)
	if err != nil {
		return host, "80"
	}
	return result, port
}

// func ioCopy(dst io.Writer, src io.Reader, fn ...func(count int)) (written int64, isSrcErr bool, err error) {
// 	buf := make([]byte, 32*1024)
// 	for {
// 		nr, er := src.Read(buf)
// 		if nr > 0 {
// 			nw, ew := dst.Write(buf[0:nr])
// 			fmt.Printf("%s\n", buf[0:nr])
// 			if nw > 0 {
// 				written += int64(nw)
// 				if len(fn) == 1 {
// 					fn[0](nw)
// 				}
// 			}
// 			if ew != nil {
// 				err = ew
// 				break
// 			}
// 			if nr != nw {
// 				err = io.ErrShortWrite
// 				break
// 			}
// 		}
// 		if er != nil {
// 			err = er
// 			isSrcErr = true
// 			break
// 		}
// 	}
// 	return written, isSrcErr, err
// }

func pipe(dst io.Writer, src io.Reader) (err error) {
	buf := make([]byte, 32*1024)
	for {
		readCount, readErr := src.Read(buf)
		if readErr != nil || readCount == 0 {
			break
		}
		fmt.Printf("write %s \n", buf[:readCount])
		writeCount, writeErr := dst.Write(buf[:readCount])
		if writeErr != nil {
			break
		}
		if readCount != writeCount {
			return errors.New("写入错误")
		}
	}
	return err
}
