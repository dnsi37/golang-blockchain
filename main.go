package main

import (
	"fmt"
	"strconv"

	"github.com/dnsi37/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("Frist Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, block := range chain.Block {
		fmt.Printf("Prev Hash : %x\n", block.PrevHash)
		fmt.Printf("Data in Block : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println(pow.Block)
	}
}
