package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

func CryptoHash(data ...[]byte) (string){
	joined := bytes.Join(data,[]byte(" "))
	hash := sha256.Sum256(joined)
	return hex.EncodeToString(hash[:])
}
