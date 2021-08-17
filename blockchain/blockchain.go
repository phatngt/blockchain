package blockchain

import (
	"bl/block"
	"bl/config"
	"bl/utils"
	"reflect"
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

func (bc *Blockchain) IsValidChain() (bool){
	//Check the chain start with genesis block
	conf := config.Genesis()
	startedChain,err := block.New(conf.Timestamp,conf.LastHash,conf.Hash,conf.Data)
	if err != nil {
		return false
	}
	if !reflect.DeepEqual(startedChain,bc.chain[0]){
		return false
	}
	//check the chain contains a block with a invalid field and a lastHash reference has changed
	for i :=1; i <len(bc.chain); i++ {
		bl := bc.chain[i]
		actualLastHash := bc.chain[i-1].Hash()
		if bl.LastHash() != actualLastHash {
			return false
		}
		validHash := utils.CryptoHash([]byte(bl.Timestamp()),[]byte(bl.LastHash()),bl.Data())
		if validHash != bl.Hash() {
			return false
		}

	}

	return true
}
