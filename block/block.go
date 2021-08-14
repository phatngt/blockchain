package block

import (
	"time"
	"crypto/sha256"
)

type Block struct {
	timestamp time
	lasthash string
	hash string
	data []string
}
var block Block
func New(timestamp time, lastHash string, hash string, data []string ) (*Block, error){
	block := &Block{
		timestamp: timestamp,
		lasthash: lastHash,
		hash: hash,
		data: data,
	}
	return block,nil
}

func (b *Block) Data() ([]string){
	return b.data
}

func (b *Block) Timestamp() string {
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

func (b *Block )MineBlock(data []string) (*Block){
	block := Block{
		timestamp: time.Now(),
		lasthash: b.lasthash,
		hash: b.hash + "abc",
		data: data
	}
	return &block
}
