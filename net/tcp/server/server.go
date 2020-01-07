package main

import (
	"flag"
	"fmt"
	"net"
)

// MaxRead 最大读取byte数
const MaxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("need host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	// init server
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "listen Tcp:")
	println("Listening to ", listener.Addr().String())
	return listener
}

// 错误处理 panic
func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

// 连接处理
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	// Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
	sendMsg(conn,"12345\r\n")
}

func sendMsg(conn net.Conn, msg string)  {
	obuf := []byte(msg)
	wrote, err := conn.Write(obuf)
	checkError(err,"Write :wrote error"+string(wrote))
}
