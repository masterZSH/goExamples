package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	//  P-224, P-256, P-384, and P-521

	curve := elliptic.P224()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	fmt.Println(privateKey, err)

	msg := "test"
	hash := sha256.Sum256([]byte(msg))
	sig, err := ecdsa.SignA	SN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)
	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)

	// 公钥
	publicKey := privateKey.Public()
	fmt.Printf("publicKey: %x\n", publicKey)

	// equal
	if k, ok := publicKey.(*ecdsa.PublicKey); ok {
		fmt.Println(k.Equal(k))
	}

}
