## rate limit




### 不用限流器

```go
file, err := os.Open("t")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1)
	start := time.Now()
	for {
		// n, err := rl.Read(buf)
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("%s\n", buf[:n])
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
```

output 
6.2086ms

### 添加限流器

```go
file, err := os.Open("t")
	if err != nil {
		panic(err)
	}

	// 1s读取10个
	rl := NewLimiterReader(file, 10)
	buf := make([]byte, 1)

	start := time.Now()
	for {
		n, err := rl.Read(buf)
		// n, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("%s\n", buf[:n])
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
```

output
4.9102269s


