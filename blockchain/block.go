package blockchain


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
  Blocks []*Block
}

func (b *Block) DeriveHash() {
  info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
  hash := sha512.Sum512(info)
  b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &Block{[]byte{}, []byte(data), prevHash}
  block.DeriveHash()
  return block
}

func (chain *BlockChain) AddBlock(data string) {
  prevBlock := chain.Blocks[len(chain.Blocks)-1]
  new := CreateBlock(data, prevBlock.Hash)
  chain.blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
  return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
  return &BlockChain{[]*Block{Genesis()}}
}
