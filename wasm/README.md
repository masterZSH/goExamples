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
goexec 'log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))'
```

http://localhost:8080/index.html



https://golangbot.com/webassembly-using-go/
