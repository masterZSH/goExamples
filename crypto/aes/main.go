package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
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

	AESCBCEncrypt()
}

func AESCBCEncrypt() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("73c86d43a9d700a253a96c85b0f6b03ac9792e0e757f869cca306bd3cba1c62b")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
}
