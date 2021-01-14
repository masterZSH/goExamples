package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

func main() {

	// 传输过来的加密数据
	t := ""

	// 申请第三方后台的填写的key
	k := ""

	aesKey := k + "="
	key, err := base64.StdEncoding.DecodeString(aesKey)
	iv := key[:16]
	deciphered, err := base64.StdEncoding.DecodeString(t)
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(deciphered, deciphered)
	decoded := delDecode(deciphered)

	buf := bytes.NewBuffer(decoded[16:20])

	var msgLen int32
	binary.Read(buf, binary.BigEndian, &msgLen)

	msgDecrypt := decoded[20 : 20+msgLen]
	id := string(decoded[20+msgLen:])
	fmt.Printf("%s,%s", key, err)

	fmt.Printf("%s,%s", msgDecrypt, id)

}

func delDecode(text []byte) []byte {
	pad := int(text[len(text)-1])

	if pad < 1 || pad > 32 {
		pad = 0
	}

	return text[:len(text)-pad]
}
