package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func Sign(params string, privateKeyStr []byte) (string, error) {
	// privateInterface, err := x509.ParsePKCS8PrivateKey(privateKeyStr)
	var privateKey *rsa.PrivateKey
	privateKey, _ = x509.ParsePKCS1PrivateKey([]byte(privateKeyStr))
	// if err != nil {

	// 	return "", err
	// }
	// privateKey, _ := privateInterface.(*rsa.PrivateKey)
	result, err := RsaSign(params, privateKey)
	if err != nil {
		return "", err
	}
	return result, nil
}

func VerifySign(raw string, signature string, publicKeyStr []byte) bool {
	publicInterface, err := x509.ParsePKIXPublicKey(publicKeyStr)
	if err != nil {
		return false
	}
	publicKey := publicInterface.(*rsa.PublicKey)

	digest := EncryptSHA([]byte(raw))

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, digest, DecryptBase64(signature))
	if err != nil {
		return false
	}
	return true
}

func RsaSign(raw string, privateKey *rsa.PrivateKey) (string, error) {
	digest := EncryptSHA([]byte(raw))

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}
