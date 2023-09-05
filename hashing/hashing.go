package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"

	"golang.org/x/crypto/sha3"
)

func CalculateSHA256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA3_256(data []byte) []byte {
	hasher := sha3.New256()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA1(data []byte) []byte {
	hasher := sha1.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateMD5(data []byte) []byte {
	hasher := md5.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func CalculateSHA512(data []byte) []byte {
	hasher := sha512.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}
