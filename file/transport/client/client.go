package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

var firstConn bool = true

// KeepAlivePeriod 长连接
const KeepAlivePeriod = time.Minute

var clients map[string]net.Conn

var (
	// MsgTypeRegister 注册消息
	MsgTypeRegister byte = '1'
	// MsgTypeFind 查询消息
	MsgTypeFind byte = '2'
)

func init() {
	clients = make(map[string]net.Conn)
}

func main() {
	go serve()
	connServer()
}

func getRandomServePort() int {
	rand.Seed(time.Now().UnixNano())
	return 1000 + rand.Intn(10000)
}

func serve() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", getRandomServePort()))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("server run %s\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// tcp conn
		if tc, ok := conn.(*net.TCPConn); ok {
			tc.SetKeepAlive(true)
			tc.SetKeepAlivePeriod(KeepAlivePeriod)
			tc.SetLinger(10)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr)
	fmt.Print(conn)
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		if n == 0 {
			continue
		}
		fmt.Printf("%s\n", buf[:n])
	}
}

func connServer() {
	scanner := bufio.NewScanner(os.Stdin)
	var conn net.Conn
	var err error
	for scanner.Scan() {
		if firstConn {
			addr := scanner.Text()
			firstConn = false
			conn, err = net.Dial("tcp", addr)
			if err != nil {
				log.Printf("conn error:%s", err)
			}
			log.Printf("conn %s server success", conn.RemoteAddr())
			continue
		}
		log.Println("please input client name")
		file := scanner.Text()
		f, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("read file err:%s\n", err)
			continue
		}
		conn.Write(f)
		break
	}
	readBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(readBuf)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s\n", readBuf[:n])
	}
}
