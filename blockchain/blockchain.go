package blockchain

import (
	"bl/block"
	"bl/config"
)

type Blockchain struct {
	chain []*block.Block
}
func (bc *Blockchain) Chain() ([]*block.Block){
	return bc.chain
}
func New() (*Blockchain){
	conf := config.Genesis()

	bl,err := block.New(conf.Timestamp,conf.LastHash,conf.Hash,conf.Data)
	if err != nil {
		return nil
	}
	blockchain := &Blockchain{chain: []*block.Block{bl}}
	return blockchain
}

func (bc *Blockchain) AddBlock(data []byte) {
	var lastCurrentBlock *block.Block
	if len(bc.chain) > 0 {
		lastCurrentBlock = bc.chain[len(bc.chain)-1]
	}
	bc.chain = append(bc.chain, lastCurrentBlock.MineBlock(data))
}
