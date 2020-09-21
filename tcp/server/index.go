package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	l, err := net.Listen("tcp", ":8989")
	delErr(err)
	for {
		conn, err := l.Accept()
		delErr(err)
		go handleConn(conn)
	}
}

func delErr(err error) {
	if err != nil {
		log.Fatalf("err:%+v\n", err)
	}
}

func handleConn(conn net.Conn) {
	// header
	var header []byte
	headerLen := 4
	for {
		tmp := make([]byte, headerLen-len(header))
		n, err := conn.Read(tmp)
		delErr(err)
		header = append(header, tmp[:n]...)
		if headerLen == len(header) {
			break
		}
	}
	var numBytesUint32 uint32
	rbuf := bytes.NewReader(header)
	err := binary.Read(rbuf, binary.LittleEndian, &numBytesUint32)
	delErr(err)
	numBytes := int(numBytesUint32)
	fmt.Printf("header : %v\n", header)
	fmt.Printf("content len : %d\n", numBytes)
	buf := make([]byte, 0)
	for {
		tmp := make([]byte, numBytes-len(buf))
		n, _ := conn.Read(tmp)
		buf = append(buf, tmp[:n]...)
		if numBytes == len(buf) {
			break
		}
	}
	fmt.Printf("content:%v\n", buf)

	k, _, _ := newC([]byte("ZSH"), []byte("Salt"))
	t, err := Decrypt(buf, k)
	fmt.Printf("decrypt %s\n", t)
}

// New generates a new key based on a passphrase and salt
func newC(passphrase []byte, usersalt []byte) (key []byte, salt []byte, err error) {
	if len(passphrase) < 1 {
		err = fmt.Errorf("need more than that for passphrase")
		return
	}
	if usersalt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		if _, err := rand.Read(salt); err != nil {
			log.Fatalf("can't get random salt: %v", err)
		}
	} else {
		salt = usersalt
	}
	key = pbkdf2.Key(passphrase, salt, 100, 32, sha256.New)
	return
}

// Encrypt will encrypt using the pre-generated key
func Encrypt(plaintext []byte, key []byte) (encrypted []byte, err error) {
	// generate a random iv each time
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	ivBytes := make([]byte, 12)
	if _, err := rand.Read(ivBytes); err != nil {
		log.Fatalf("can't initialize crypto: %v", err)
	}
	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	encrypted = aesgcm.Seal(nil, ivBytes, plaintext, nil)
	encrypted = append(ivBytes, encrypted...)
	return
}

// Decrypt using the pre-generated key
func Decrypt(encrypted []byte, key []byte) (plaintext []byte, err error) {
	if len(encrypted) < 13 {
		err = fmt.Errorf("incorrect passphrase")
		return
	}
	b, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return
	}
	plaintext, err = aesgcm.Open(nil, encrypted[:12], encrypted[12:], nil)
	return
}
