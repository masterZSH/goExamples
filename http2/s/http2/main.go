package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync/atomic"
	"time"

	"golang.org/x/net/http2"
)

const cipher_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 uint16 = 0xC02F

type fakeTLSConn struct {
	net.Conn
}

func (c *fakeTLSConn) ConnectionState() tls.ConnectionState {
	return tls.ConnectionState{
		Version:     tls.VersionTLS12,
		CipherSuite: cipher_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}
}

func main() {
	// var conn http2.ClientConn
	// fmt.Println(conn.CanTakeNewRequest())
	l := newLocalListener()
	defer l.Close()
	h2Server := &http2.Server{}
	go func() {
		req, err := http.NewRequest("GET", "http://"+l.Addr().String()+"/foobar", nil)
		fmt.Println(req, err)
		var gotConnCnt int32
		trace := &httptrace.ClientTrace{
			GotConn: func(connInfo httptrace.GotConnInfo) {
				if !connInfo.Reused {
					atomic.AddInt32(&gotConnCnt, 1)
				}
			},
		}
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		tr := &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		}
		res, err := tr.RoundTrip(req)

		fmt.Println(res, err)
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		h2Server.ServeConn(&fakeTLSConn{conn}, &http2.ServeConnOpts{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)
		})})
	}
}

func newLocalListener() net.Listener {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err == nil {
		fmt.Printf("serve listen %s\n", ln.Addr())
		return ln
	}
	ln, err = net.Listen("tcp6", "[::1]:0")
	if err != nil {
		log.Fatal(err)
	}
	return ln
}

func serve() {
	h := handler{}
	s := &http.Server{
		Addr:           ":8081",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v, http: %v", r.URL.Path, r.TLS == nil)

}
