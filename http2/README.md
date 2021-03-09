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


GODEBUG=http2client=0  # disable HTTP/2 client support
GODEBUG=http2server=0  # disable HTTP/2 server support
GODEBUG=http2debug=1   # enable verbose HTTP/2 debug logs
GODEBUG=http2debug=2   # ... even more verbose, with frame dumps