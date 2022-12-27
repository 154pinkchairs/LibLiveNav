package main

import (
	"fmt"
	"strconv"

	"github.com/154pinkchairs/LibLiveNav/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("1st block after Genesis")
	chain.AddBlock("2nd block after Genesis")
	chain.AddBlock("3rd block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
