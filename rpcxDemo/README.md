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



## etcd  v3.4

1. 新增role

新增角色root
```bash
./etcdctl role add root
```

2. 授权

给root角色授权读写权限 键空间 /zsh/*(*通配符)
```bash
./etcdctl role grant-permission root readwrite '/zsh/*'
```

3. 新增用户

增加用户zsh密码123456
```bash
./etcdctl user add zsh:123456
```

增加用户root密码root
```bash
./etcdctl user add root:root
```

4. 用户角色分配

给zsh用户分配root角色
```bash
./etcdctl user grant-role zsh root
```

给root用户分配root角色
```bash
./etcdctl user grant-role root root
```

5. 启动授权认证

```bash
./etcdctl auth enable
```
