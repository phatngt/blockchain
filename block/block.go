package block

import (
	"bl/utils"
	"time"
)

type Block struct {
	timestamp string
	lasthash string
	hash string
	data []byte
}
var block Block
func New(timestamp string, lastHash string, hash string, data []byte ) (*Block, error){
	block := &Block{
		timestamp: timestamp,
		lasthash: lastHash,
		hash: utils.CryptoHash([]byte(timestamp),[]byte(lastHash),data),
		data: data,
	}
	return block,nil
}

func (b *Block) Data() ([]byte){
	return b.data
}

func (b *Block) SetData(data []byte){
	b.data = data
}

func (b *Block) Timestamp() string {
	return b.timestamp
}

func (b *Block) SetTimestamp(timestamp string){
	b.timestamp = timestamp
}

func (b *Block) Hash() string {
	return b.hash
}

func (b *Block) SetHash(hash string) {
	b.hash = hash
}

func (b *Block) LastHash() string {
	return b.lasthash
}

func (b *Block) SetLastHast(lasthash string){
	b.lasthash = lasthash
}

func GetBlock() (*Block){
	return &block;
}

func (b *Block )MineBlock(data []byte) (*Block){
	timestamp := time.Now().Format(time.RFC1123)
	bl := Block{
		timestamp: timestamp,
		lasthash: b.hash,
		hash: utils.CryptoHash([]byte(timestamp),[]byte(b.hash),data),
		data: data,
	}
	return &bl
}
