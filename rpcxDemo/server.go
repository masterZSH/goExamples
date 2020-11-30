package main

import (
	"context"
	"flag"
	"log"
	"net"

	ah "github.com/masterZSH/rservices/auth"
	"github.com/masterZSH/rservices/user"

	"time"

	"github.com/docker/libkv/store"

	graphite "github.com/cyberdelia/go-metrics-graphite"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
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
	s.Register(new(user.User), "")
	// auth
	s.AuthFunc = auth
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
		log.Fatal(err)
	}
	s.Plugins.Add(r)

	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	// 限流 30s 1000
	limit := serverplugin.NewRateLimitingPlugin(30*time.Second, 10000)
	s.Plugins.Add(limit)

}

func auth(ctx context.Context, req *protocol.Message, token string) error {
	err := validateToken(token)
	if err != nil {
		return ah.ErrInvalidToken
	}
	return nil
}

func validateToken(tokenString string) error {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.ConnectTimeout = 10 * time.Second
	// Auth
	cg := &store.Config{
		Username: "zsh",
		Password: "123456",
	}
	d := client.NewEtcdV3Discovery(*basePath, "Validate", []string{*etcdAddr}, cg)
	xclient := client.NewXClient("Validate", client.Failfast, client.RandomSelect, d, option)
	defer xclient.Close()
	args := &ah.ValidateArgs{
		Token: tokenString,
	}
	reply := &ah.ValidateReply{}
	// 同步 异步使用Go
	err := xclient.Call(context.Background(), "ValidateToken", args, reply)
	if err != nil {
		return err
	}
	return err
}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	// 配置graphite地址
	addr, _ := net.ResolveTCPAddr("tcp", "111.229.204.111:2003")

	// 监控节点名称 rpcx.services.host + 服务器ip地址
	go graphite.Graphite(metrics.DefaultRegistry, 1e9, "rpcx.services.host.127_0_0_1", addr)
}
