package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

const NonceLength = 12

// key len 16,      24,      or 32
//         AES-128, AES-192, or AES-256.
func Encrypt(key []byte, data []byte) ([]byte, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, NonceLength)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Prefix sealed data with nonce for decryption - see below.
	return append(nonce, gcm.Seal(nil, nonce, data, nil)...), nil
}

// Decrypt uses key to decrypt the given data.
func Decrypt(key []byte, data []byte) ([]byte, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Use prefixed nonce from above.
	return gcm.Open(nil, data[:NonceLength], data[NonceLength:], nil)
}

func main() {
	key := []byte("1231231231231231")
	data := []byte("456")
	l, err := Encrypt(key, data)
	fmt.Println(l)
	fmt.Println(err)
	t, err := Decrypt(key, l)
	fmt.Println(err)
	fmt.Printf("%s\n", t)
}
