package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/task"
)

func Server() {
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2666", cfg)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

type readerOnly struct {
	io.Reader
}

func handleConn(conn net.Conn) {
	for {
		reader := bufio.NewReaderSize(readerOnly{conn}, 2048)
		r, err := http.ReadRequest(reader)
		fmt.Print(err)
		if err != nil {
			continue
		}

		defaultPort := 80
		if strings.EqualFold(r.URL.Scheme, "https") {
			defaultPort = 443
		}
		host := r.Host
		if host == "" {
			host = r.URL.Host
		}
		if strings.EqualFold(r.Method, "CONNECT") {
			handleConnect(conn)
		} else {

		}

	}

}

func handleConnect(conn net.Conn) {
	_, err := conn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func handleHttp(request *http.Request, conn net.Conn) error {
	if request.URL.Host == "" {
		response := &http.Response{
			Status:        "Bad Request",
			StatusCode:    400,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Header:        http.Header(make(map[string][]string)),
			Body:          nil,
			ContentLength: 0,
			Close:         true,
		}
		response.Header.Set("Proxy-Connection", "close")
		response.Header.Set("Connection", "close")
		return response.Write(conn)
	}

	if len(request.URL.Host) > 0 {
		request.Host = request.URL.Host
	}

	removeHeader(request.Header)

	// Prevent UA from being set to golang's default ones
	if request.Header.Get("User-Agent") == "" {
		request.Header.Set("User-Agent", "")
	}

	for key := range request.Header {
		value := request.Header.Get(key)
	}

	requestDone := func() error {
		request.Header.Set("Connection", "close")

		requestWriter := buf.NewBufferedWriter(link.Writer)
		common.Must(requestWriter.SetBuffered(false))
		if err := request.Write(requestWriter); err != nil {
		}
		return nil
	}

	responseDone := func() error {
		responseReader := bufio.NewReaderSize(&buf.BufferedReader{Reader: link.Reader}, buf.Size)
		response, err := http.ReadResponse(responseReader, request)
		if err == nil {
			if response.ContentLength >= 0 {
				response.Header.Set("Proxy-Connection", "keep-alive")
				response.Header.Set("Connection", "keep-alive")
				response.Header.Set("Keep-Alive", "timeout=4")
				response.Close = false
			} else {
				response.Close = true
				result = nil
			}
		} else {
			response = &http.Response{
				Status:        "Service Unavailable",
				StatusCode:    503,
				Proto:         "HTTP/1.1",
				ProtoMajor:    1,
				ProtoMinor:    1,
				Header:        http.Header(make(map[string][]string)),
				Body:          nil,
				ContentLength: 0,
				Close:         true,
			}
			response.Header.Set("Connection", "close")
			response.Header.Set("Proxy-Connection", "close")
		}
		if err := response.Write(writer); err != nil {
			return newError("failed to write response").Base(err).AtWarning()
		}
		return nil
	}

	if err := task.Run(ctx, requestDone, responseDone); err != nil {
		common.Interrupt(link.Reader)
		common.Interrupt(link.Writer)
		return newError("connection ends").Base(err)
	}

	return result
}

func removeHeader(header http.Header) {
	header.Del("Proxy-Connection")
	header.Del("Proxy-Authenticate")
	header.Del("Proxy-Authorization")
	header.Del("TE")
	header.Del("Trailers")
	header.Del("Transfer-Encoding")
	header.Del("Upgrade")

	connections := header.Get("Connection")
	header.Del("Connection")
	if connections == "" {
		return
	}
	for _, h := range strings.Split(connections, ",") {
		header.Del(strings.TrimSpace(h))
	}
}
