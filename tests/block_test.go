package tests

import (
	"bl/block"
	"testing"
	"time"
)

func TestBlock(t *testing.T) {
	bl, err := block.New(time.Now(),"foo-lasthash","foo-hash",[]string{"foo-data"})
	got := bl.Data();
	want := []string{"foo-data"}
	if got[0] != want[0] {
		t.Errorf("got: %q not equal wanted: %q",got, want)
	}
	if err != nil {
		t.Errorf("got: %q not equal wanted: %v",err, nil)
	}

}

func TestMine(t *testing.T){
	lastBlock := block.GetBlock()
	data := []string{"new-data"}
	mineBlock := lastBlock.MineBlock(data)

	got0 := mineBlock.Data()[0]
	want0 := data[0]
	if got0 != want0 {
		t.Errorf("got: %q not equal wanted: %q",got0,want0)
	}

	got1 := mineBlock.LastHash()
	want1 := lastBlock.LastHash()
	if got1 != want1 {
		t.Errorf("got: %q not equal wanted: %q",got1,want1)
	}



}
