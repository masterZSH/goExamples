package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	const size = 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		panic(err)
	}
	secretMessage := []byte("send reinforcements, we're going to advance")
	label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader
	hash := sha256.New()

	// public key 加密
	ciphertext, err := rsa.EncryptOAEP(hash, rng, &privateKey.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// private key 解密
	pureText, err := rsa.DecryptOAEP(hash, rng, privateKey, ciphertext, label)
	if err != nil {
		return
	}
	fmt.Printf("pureText: %s\n", pureText)

}
