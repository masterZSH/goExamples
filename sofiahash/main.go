package main

import (
	"crypto/md5"
	"log"
)

func SofiaHash(password string) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	sofia := make([]byte, 0, 8)
	hash := md5.Sum([]byte(password))
	for i := 0; i < md5.Size; i += 2 {
		j := uint16(hash[i]) + uint16(hash[i+1])
		sofia = append(sofia, chars[j%62])
	}

	return string(sofia)
}

func main() {
	cipherText := SofiaHash("123456")
	log.Println(cipherText)
}
