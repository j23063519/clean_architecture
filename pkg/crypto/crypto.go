package crypto

import (
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
)

func AesEncrypt(val string) string {
	return crypto.
		FromString(val).
		SetIv("").
		SetKey("").
		Aes().
		CBC().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
}

func AesDecrypt(val string) string {
	return crypto.
		FromBase64String(val).
		SetIv("").
		SetKey("").
		Aes().
		CBC().
		PKCS7Padding().
		Decrypt().
		ToString()
}

func DesEncrypt(val string) string {
	return crypto.
		FromString(val).
		SetIv("").
		SetKey("").
		Des().
		CBC().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
}

func DesDecrypt(val string) string {
	return crypto.
		FromBase64String(val).
		SetIv("").
		SetKey("").
		Des().
		CBC().
		PKCS7Padding().
		Decrypt().
		ToString()
}

func TripleDesEncrypt(val string) string {
	return crypto.
		FromString(val).
		SetIv("").
		SetKey("").
		TripleDes().
		CBC().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
}

func TripleDesDecrypt(val string) string {
	return crypto.
		FromBase64String(val).
		SetIv("").
		SetKey("").
		TripleDes().
		CBC().
		PKCS7Padding().
		Decrypt().
		ToString()
}
