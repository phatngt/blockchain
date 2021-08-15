package main

import (
	"bl/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.New()
	fmt.Println(bc.Chain()[0])
}
