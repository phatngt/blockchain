package tests

import (
	"bl/blockchain"
	"testing"
)

// func TestBlockChain(t *testing.T) {
// 	//describe: Block chain
// 	blockChain := blockchain.New()
// 	blockInstance := blockChain.Chain()[0]

// 	//it: Contains a `chain` Array instance
// 	var typeOfChain []*block.Block
// 	got0 := reflect.TypeOf(blockChain.Chain())
// 	want0 := reflect.TypeOf(typeOfChain)
// 	if got0 != want0 {
// 		t.Errorf("Got: %q not equal to wanted: %q",got0,want0)
// 	}

// 	//it: start with genesis block
// 	lastHash := utils.CryptoHash([]byte("-----"))
// 	got1 := blockInstance.LastHash()
// 	want1 := lastHash

// 	if got1 != want1 {
// 		t.Errorf("Got: %q not equal to wanted: %q",got1,want1)
// 	}

// 	//it: add new a block to the chain
// 	var data []byte = []byte("new-data")
// 	blockChain.AddBlock(data)
// 	chain := blockChain.Chain()
// 	got2 := chain[len(chain)-1].Data()
// 	want2 := data
// 	if !bytes.Equal(got2,want2) {
// 		t.Errorf("Got: %q not equal to wanted: %q",got2,want2)
// 	}

// 	//it: add multiple blocks to the chain
// 	data = []byte("one-data")
// 	blockChain.AddBlock(data)
// 	data = []byte("two-data")
// 	blockChain.AddBlock(data)
// 	chain = blockChain.Chain()
// 	got3 := len(chain)
// 	want3 := 4
// 	if got3 != want3 {
// 		t.Errorf("Got: %q not equal to wanted: %q",got3,want3)
// 	}
// 	got4 := chain[2].Data()
// 	want4 := []byte("one-data")
// 	if !bytes.Equal(got4,want4) {
// 		t.Errorf("Got: %q not equal to wanted: %q",got4,want4)
// 	}
// }

func TestIsValidChain(t *testing.T) {
	//Describe: Invalid chain
	blockChain := blockchain.New()
	// conf := config.Genesis()

	//it-combine when the chain starts with the genesis block and has multiple blocks
	blockChain.AddBlock([]byte("one-data"))
	blockChain.AddBlock([]byte("two-data"))

	//it: and the chain does not contain any invalid block
	got1 := blockChain.IsValidChain()
	want1 := true
	if got1 != want1 {
		t.Errorf("Got: %t not equal wanted: %t",got1,want1)
	}

	// //it: when the chain does not start with the genesis block
	// var err error
	// blockChain.Chain()[0],err = block.New(conf.Timestamp,conf.LastHash,conf.Hash,[]byte("fake-data"))
	// if err != nil{
	// 	t.Errorf("Error when create new block: %q",err)
	// 	return
	// }

	// got0 := blockChain.IsValidChain()
	// want0 := false
	// if got0 != want0 {
	// 	t.Errorf("Got: %t not equal wanted: %t",got0,want0)
	// }

	// //it: and the chain contains a block with a invalid field
	// blockChain.Chain()[2].SetData([]byte("some-bad-and-evil-data"))
	// got2 := blockChain.IsValidChain()
	// want2 := false
	// if got2 != want2 {
	// 	t.Errorf("Got: %t not equal wanted: %t",got2,want2)
	// }
	// blockChain.Chain()[2].SetData(([]byte("two-data")))

	// //it: and a lastHash reference has changed
	// blockChain.Chain()[2].SetLastHast("broken-hash")
	// got3 := blockChain.IsValidChain()
	// want3 := false
	// if got3 != want3 {
	// 	t.Errorf("Got: %t not equal wanted: %t",got3,want3)
	// }


	//it:


}
