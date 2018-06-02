package main

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)

const (
	RSAPublicKey = "-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCKFctVrhfF3m2Kes0FBL/JFeO\n" +
		"cmNg9eJz8k/hQy1kadD+XFUpluRqa//Uxp2s9W2qE0EoUCu59ugcf/p7lGuL99Uo\n" +
		"SGmQEynkBvZct+/M40L0E0rZ4BVgzLOJmIbXMp0J4PnPcb6VLZvxazGcmSfjauC7\n" +
		"F3yWYqUbZd/HCBtawwIDAQAB\n" +
		"-----END PUBLIC KEY-----"
)

func RSAEncryptString(s string) string {
	byteString := []byte(s)
	block, _ := pem.Decode([]byte(RSAPublicKey))
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	encBytes, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, byteString)
	return string(encBytes)
}

func main() {
	stringToEncrypt := "this_is_my_string"
	encryptedString := RSAEncryptString(stringToEncrypt)
	doWork(encryptedString)
}

func doWork(s string) {

}
