package utils

import (
	"crypto/sha256"

	base58 "github.com/btcsuite/btcutil/base58"
)

func Sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func Base58Encoded(bytes []byte) string {
	encoded := base58.Encode(bytes)
	return encoded
}
