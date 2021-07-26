package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

//
// HTTP/1.x
//

var (
	doneChan chan interface{}
)

var (
	ErrServerClosed = errors.New("服务已关闭")
)

func init() {
	doneChan = make(chan interface{})
}

func listenAndServe(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	return Serve(l)
}

func Serve(l net.Listener) (err error) {
	defer l.Close()
	// ctx := context.Background()
	for {
		rw, err := l.Accept()
		if err != nil {
			select {
			case <-doneChan:
				return ErrServerClosed
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				continue
			}
			return err
		}
		go serveConn(rw)
	}

}

type connReader struct {
	conn   net.Conn
	inRead bool
	mu     sync.Mutex
}

func (r *connReader) lock() {
	r.mu.Lock()
}

func (r *connReader) unlock() {
	r.mu.Unlock()
}

func (r *connReader) Read(p []byte) (n int, err error) {
	r.lock()
	if r.inRead {
		r.unlock()
		panic("同时读取错误")
	}
	if len(p) == 0 {
		r.unlock()
		return 0, nil
	}
	r.inRead = true
	r.unlock()
	n, err = r.conn.Read(p)
	r.lock()
	r.inRead = false
	r.unlock()
	return n, err
}

func serveConn(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	// tls connection
	if tlsConn, ok := conn.(*tls.Conn); ok {
		fmt.Println(tlsConn)
	}

	// set read deadline
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// set write deadline
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	//
	//
	// read large data
	buf := make([]byte, 0, 4<<10)
	tmp := make([]byte, 256)
	for {
		// read from conn
		// r := &connReader{conn: conn}
		// buf := bufio.NewReader(r)
		// r.Read(p)
		p := make([]byte, 5)
		n, er := conn.Read(p)
		atEOF := errors.Is(er, io.EOF)
		if atEOF {
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	fmt.Printf("total size:%d\n", len(buf))
}

func main() {
	go func() {
		listenAndServe(":9823")
	}()

	conn, err := net.Dial("tcp", "localhost:9823")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write(s)
	select {}
}
