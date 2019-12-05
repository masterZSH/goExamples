package main

import "os"

// Count byte count
const Count int = 512

func main() {
	var bytes [Count]byte
	f, _ := os.Open("test.text")
	defer f.Close()
	for {
		i, _ := f.Read(bytes[:])
		os.Stdout.Write(bytes[0:i])
	}

}
