package config

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Config struct {
	Timestamp time.Time
	LastHash string
	Hash string
	Data []byte
}

func Genesis() ( *Config){
	initHash := sha256.Sum256([]byte("init-block"))
	lastHash := sha256.Sum256([]byte("-----"))
	data := []byte("")
	config := &Config{Timestamp: time.Now(), LastHash: hex.EncodeToString(lastHash[:]), Hash: hex.EncodeToString(initHash[:]),Data: data}
	return config
}
