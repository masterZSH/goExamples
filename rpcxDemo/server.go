package main

import (
	"flag"
	"main/rpc/services/user"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8973", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	//s.RegisterName("User", new(user.User), "")
	s.Register(new(user.User), "")
	s.Serve("tcp", *addr)
}
