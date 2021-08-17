package config

import (
	"bl/utils"
	"time"
)

type Config struct {
	Timestamp string
	LastHash string
	Hash string
	Data []byte
}

func Genesis() ( *Config){

	date := "Tue, 17 Aug 2021 08:44:22 +07"
	timestamp, _ := time.Parse(time.RFC1123, date)
	lastHash := utils.CryptoHash([]byte("-----"))
	data := []byte("")
	initHash := utils.CryptoHash([]byte(timestamp.String()),[]byte(lastHash),data)
	config := &Config{Timestamp: timestamp.String(), LastHash:lastHash, Hash: initHash,Data: data}


	return config
}
