package models

import "crypto/sha256"

type EncryptedData []byte

func Encrypt(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

type Email []byte
