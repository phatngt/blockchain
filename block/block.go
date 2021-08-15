package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	timestamp time.Time
	lasthash string
	hash string
	data []byte
}
var block Block
func New(timestamp time.Time, lastHash string, hash string, data []byte ) (*Block, error){
	block := &Block{
		timestamp: timestamp,
		lasthash: lastHash,
		hash: hash,
		data: data,
	}
	return block,nil
}

func (b *Block) Data() ([]byte){
	return b.data
}

func (b *Block) Timestamp() time.Time {
	return b.timestamp
}

func (b *Block) Hash() string {
	return b.hash
}

func (b *Block) LastHash() string {
	return b.lasthash
}

func GetBlock() (*Block){
	return &block;
}

func (b *Block )MineBlock(data []byte) (*Block){
	hash := sha256.Sum256(data)
	block := Block{
		timestamp: time.Now(),
		lasthash: b.lasthash,
		hash: hex.EncodeToString(hash[:]),
		data: data,
	}
	return &block
}
