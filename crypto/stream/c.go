package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	encrypted, _ := hex.DecodeString("cf0495cc6f75dafc23948538e79904a9")
	bReader := bytes.NewReader(encrypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	reader := &cipher.StreamReader{S: stream, R: bReader}
	// Copy the input to the output stream, decrypting as we go.
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the output.

	mReader := bytes.NewReader([]byte("some secret text"))
	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	if _, err := io.Copy(writer, mReader); err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", out.Bytes())

}
