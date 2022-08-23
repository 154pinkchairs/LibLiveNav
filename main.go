package main

import (
  "fmt"
  "github.com/154pinkchairs/LibLiveNav/blockchain"
)

func main()  {
  chain := InitBlockChain()

  chain.AddBlock("1st block after Genesis")
  chain.AddBlock("2nd block after Genesis")
  chain.AddBlock("3rd block after Genesis")

  for _, block := range chain.blocks {
    fmt.Printf("Previous hash: %x\n", block.PrevHash)
    fmt.Printf("Data in block: %x\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
  }
}
