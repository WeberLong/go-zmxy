package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func DecryptRSA(data string, privateKeyStr []byte) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyStr)
	if err != nil {
		return "", err
	}

	k := (privateKey.N.BitLen() + 7) / 8
	decrypted := make([]byte, 0, len(cipherText))
	for i := 0; i < len(cipherText); i += k {
		if i+k < len(cipherText) {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText[i:i+k])
			if err1 != nil {
				return "", err
			}
			decrypted = append(decrypted, partial...)
		} else {
			partial, err1 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText[i:])
			if err1 != nil {
				return "", err
			}
			decrypted = append(decrypted, partial...)
		}
	}

	return string(decrypted), err
}

func DecryptBase64(data string) []byte {
	decrypted, _ := base64.StdEncoding.DecodeString(data)
	return decrypted
}
