package tests

import (
	"bl/block"
	"bl/blockchain"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestBlockChain(t *testing.T) {
	//describe: Block chain
	blockChain := blockchain.New()
	blockInstance := blockChain.Chain()[0]

	//it: Contains a `chain` Array instance
	var typeOfChain []*block.Block
	got0 := reflect.TypeOf(blockChain.Chain())
	want0 := reflect.TypeOf(typeOfChain)
	if got0 != want0 {
		t.Errorf("Got: %q not equal to wanted: %q",got0,want0)
	}

	//it: start with genesis block
	hashRaw := sha256.Sum256([]byte("-----"))
	lastHash := hex.EncodeToString(hashRaw[:])
	got1 := blockInstance.LastHash()
	want1 := lastHash

	if got1 != want1 {
		t.Errorf("Got: %q not equal to wanted: %q",got1,want1)
	}

	hashRaw = sha256.Sum256([]byte("init-block"))
	hash := hex.EncodeToString(hashRaw[:])
	got1 = blockInstance.Hash()
	want1 = hash
	if got1 != want1 {
		t.Errorf("Got: %q not equal to wanted: %q",got1,want1)
	}
	//it: add new a block to the chain
	var data []byte = []byte("new-data")
	blockChain.AddBlock(data)
	chain := blockChain.Chain()
	got2 := chain[len(chain)-1].Data()
	want2 := data
	if !bytes.Equal(got2,want2) {
		t.Errorf("Got: %q not equal to wanted: %q",got2,want2)
	}
}

func TestIsValidChain(t *testing.T) {
	//Describe: Invalid chain

	//it: when the chain does not start with the genesis block

	//it: when the chain starts with the genesis block and has multiple blocks



}
