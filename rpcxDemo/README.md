```
cd rpcxDemo 

go run server.go

cd client

go run client.go

```




## 错误记录


1. undefined: grpc.Address

etcd和grpc版本冲突，在go.mod中添加
```
replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
```

2. etcd3版本

```
// 使用etcd3版本时候的服务发现
r := &serverplugin.EtcdRegisterPlugin

// 替换为
r := &serverplugin.EtcdV3RegisterPlugin

```