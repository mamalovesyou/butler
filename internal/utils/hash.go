package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// SHA256 hashes using sha256 algorithm
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
