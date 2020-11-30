## 项目

### 网关 gateway
将http请求分发到微服务

### auth 认证服务 

客户端获取授权访问服务节点，验证授权信息是否有效

### client 客户端

通过etcd3查询服务并请求服务

### server.go 主服务入口

注册获取用户信息...服务




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

6. 查询所有key

```bash
./etcdctl --user root:root get --from-key ''
```

启动3个节点 无tls版本
```

etcd --name infra1 --listen-client-urls http://localhost:2379 --advertise-client-urls http://127.0.0.1:2379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr  


etcd --name infra2  --listen-client-urls http://localhost:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr 


etcd --name infra3 --listen-client-urls http://localhost:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls  http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
```



## graphite 监控服务

```
docker run -d\
 --name graphite\
 --restart=always\
 -p 8110:80\
 -p 2003-2004:2003-2004\
 -p 2023-2024:2023-2024\
 -p 8125:8125/udp\
 -p 8126:8126\
 graphiteapp/graphite-statsd
```

web端口8110
默认账户 root:root


更新时区

进入容器 /opt/graphite/webapp/graphite/ 修改settings.py/local_settings.py TIME_ZONE = 'Asia/Shanghai'
