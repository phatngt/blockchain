package tests

import (
	"bl/block"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {
	bl, err := block.New(time.Now(),"foo-lasthash","foo-hash",[]byte("foo-data"))
	got := bl.Data();
	want := []byte("foo-data")
	if got[0] != want[0] {
		t.Errorf("got: %q not equal wanted: %q",got, want)
	}
	if err != nil {
		t.Errorf("got: %q not equal wanted: %v",err, nil)
	}

}

func TestMine(t *testing.T){
	lastBlock := block.GetBlock()
	data := []byte("new-data")
	mineBlock := lastBlock.MineBlock(data)

	got0 := mineBlock.Data()
	want0 := data
	if !bytes.Equal(got0,want0) {
		t.Errorf("got: %q not equal wanted: %q",got0,want0)
	}

	got1 := mineBlock.LastHash()
	want1 := lastBlock.LastHash()
	if got1 != want1 {
		t.Errorf("got: %q not equal wanted: %q",got1,want1)
	}

	hash := sha256.Sum256(data)
	got2 := mineBlock.Hash()
	want2 :=  hex.EncodeToString(hash[:])
	if got2 != want2 {
		t.Errorf("got: %q not equal wanted: %q",got1,want1)
	}

}
