package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	// "io/ioutil"
	// "os"
)

func GetPemKeyFile(pemKey string) ([]byte, error) {
	// f, err := os.Open(path)
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()
	// pemKey, err := ioutil.ReadAll(f)
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("PEM decode failed")
	}
	return block.Bytes, nil
}

func EncryptRSA(data, publicKeyStr []byte) ([]byte, error) {
	pubInterface, err := x509.ParsePKIXPublicKey(publicKeyStr)
	if err != nil {
		return nil, err
	}
	pubicKey := pubInterface.(*rsa.PublicKey)
	var encryptedDataBuffer bytes.Buffer
	for datalens := len(data); datalens > 0; datalens = len(data) {
		maxBlock := 117
		if datalens < maxBlock {
			maxBlock = datalens
		}
		input := data[:maxBlock]
		data = data[maxBlock:]
		encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubicKey, input)
		if err != nil {
			return nil, err
		}
		encryptedDataBuffer.Write(encryptedData)
	}
	return encryptedDataBuffer.Bytes(), nil
}

func EncryptSHA(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
