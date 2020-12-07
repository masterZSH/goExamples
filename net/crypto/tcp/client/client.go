package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
)

func main() {
	pem, err := ioutil.ReadFile("cert/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(pem))
	if !ok {
		log.Fatal(errors.New("parse pem error\n"))
	}
	conn, err := tls.Dial("tcp", ":8080", &tls.Config{
		RootCAs:    roots,
		ServerName: "localhost",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
}
