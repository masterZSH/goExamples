package main

import (
	"flag"
	"log"
	"time"

	"github.com/masterZSH/rservices/auth"

	"github.com/docker/libkv/store"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr     = flag.String("addr", "localhost:9999", "auth server address")
	etcdAddr = flag.String("etcd", "127.0.0.1:2379", "etcd address")
	basePath = flag.String("basePath", "zsh", "base path")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	addRegistryPlugin(s)
	s.Register(new(auth.Auth), "")
	s.Register(new(auth.Validate), "")
	s.Serve("tcp", *addr)
}

func addRegistryPlugin(s *server.Server) {

	// 注意etcd版本
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
		Options: &store.Config{
			Username: "zsh",
			Password: "123456",
		},
	}
	err := r.Start()
	if err != nil {
		return
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
