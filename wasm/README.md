## WASM

build
```
GOOS=js GOARCH=wasm go build -o main.wasm
```

copy wasm_exec
```
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```

start
```
goexec 'http.ListenAndServe(`:8080`, http.FileServertp.Dir(`.`)))'
```

http://localhost:8080/index.html

