module main

go 1.14

require (
	github.com/cyberdelia/go-metrics-graphite v0.0.0-20161219230853-39f87cc3b432
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/libkv v0.2.1
	github.com/gin-gonic/gin v1.6.3
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/masterZSH/rservices v0.0.0-20201127085618-f80e43867c9b
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-gateway v0.0.0-20200521025828-a39934d3752d
	github.com/smallnest/rpcx v0.0.0-20201027145221-c31b15be63d4
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
