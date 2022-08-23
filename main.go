package main

import (
  "bytes"
  "crypto/sha512"
)

type Block struct {
  Hash     []byte
  Data     []byte
  PrevHash []byte
}

type BlockChain struct {
  blocks []*Block
}

func (b *Block) DeriveHash() {
  info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
  hash := sha512.Sum512(info)
  b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &{[]byte{}, []byte(data), prevHash}
  block.DeriveHash()
  return block
}

func (chain *BlockChain) AddBlock(data string) {
  prevBlock := chain.blocks[len(chain.blocks)-1]
  new := CreateBlock(data, prevBlock.Hash)
  chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
  return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
  return &BlockChain{[]*Block{Genesis()}}
}

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
