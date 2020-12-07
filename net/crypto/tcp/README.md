## 证书生成


创建密钥key 
```bash
openssl genrsa -out key.pem 2048
```

生成证书请求文件
```bash
openssl req -new -key key.pem -out cert.csr
```

生成证书文件  cert.pem 
```bash
openssl req -new -x509 -key key.pem -out cert.pem -days 1095
```


2020/12/07 10:12:41 x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0       
exit status 1
```bash
export GODEBUG=x509ignoreCN=0
```


### todo
opensll san证书