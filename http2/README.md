## http2

证书生成
```bash
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
```

server
```bash
cd http2 && go run server/index.go
```

client 
```bash
cd http && go run client/index.go
```


## 其他
http2的更多高级特性后面继续练习
