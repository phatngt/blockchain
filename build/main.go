package main

import (
	"bl/blockchain"
	"bl/utils"
	"fmt"
)

func main() {
	bc := blockchain.New()
	fmt.Println(bc.Chain()[0])
	result := utils.CryptoHash([]byte("test1"),[]byte("test2"),[]byte("test3"))
	fmt.Printf("Hash: %q",result)
}
