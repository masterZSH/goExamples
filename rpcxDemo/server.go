package main

import (
	"flag"
	"log"
	"main/rpc/services/user"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8973", "server address")

	etcdAddr = flag.String("etcd", "127.0.0.1:2379", "etcd address")

	basePath = flag.String("basePath", "zsh", "base path")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	addRegistryPlugin(s)
	//s.RegisterName("User", new(user.User), "")
	s.Register(new(user.User), "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
