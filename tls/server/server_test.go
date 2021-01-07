package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	go Server()
	time.Sleep(3 * time.Second)

	conn, err := tls.Dial("tcp", "localhost:2666", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	// test 123
	test := "CONNECT www.baidu.com:443 HTTP/1.1\r\nHost: www.baidu.com:443\r\n\r\n"
	// go handleCConn(conn)
	n, err := conn.Write([]byte(test))
	assert.Nil(t, err)
	fmt.Print(n)
	for {

	}
}

func handleCConn(conn net.Conn) {
	rd := bufio.NewReader(conn)
	s, err := rd.ReadString('\n')
	fmt.Print(s, err)
}
