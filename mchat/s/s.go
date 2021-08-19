package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/masterZSH/mino"

	"github.com/schollz/peerdiscovery"
)

type Messager interface {
	Send(data []byte) (int, error)
}

type Message struct {
	Type int
	Data []byte
}

type HConn struct {
	conn         net.Conn
	FirstConnect bool
}

func newConn(conn net.Conn) *HConn {
	return &HConn{
		conn:         conn,
		FirstConnect: true,
	}

}

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     1,
		AllowSelf: true,
		Payload:   []byte("chat:" + port),
	})
	for _, d := range discoveries {
		fmt.Printf("discovered '%s'\n", d.Address)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept tcp error")
		}

		go handleConn(newConn(conn))
	}

}

func handleConn(conn *HConn) {

	// send msg
	m := Message{
		Type: 1,
		Data: []byte("123-456"),
	}
	mt, _ := json.Marshal(m)
	mt = append(mt, '\n')
	conn.conn.Write(mt)
	conn.FirstConnect = false
	str, err := bufio.NewReader(conn.conn).ReadBytes('\n')
	if err != nil {
		return
	}
	fmt.Println("Message Received: ", string(str))

	var m1 Message
	json.Unmarshal(str, &m1)

	k, _ := mino.NewKey([]byte("123"), []byte("456"))
	content, _ := k.Decrypt(m1.Data)
	fmt.Printf("%s\n", content)
}
