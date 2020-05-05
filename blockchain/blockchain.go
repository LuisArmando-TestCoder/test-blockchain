package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash	 []byte
	Data	 string
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (block *Block) InsertHash() {
	info := bytes.Join([][]byte{[]byte(block.Data), block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash}
	block.InsertHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) *Block {
	prevBlock := chain.Blocks[len(chain.Blocks) - 1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
	return new
}

func Genesis() *Block {
	return CreateBlock("The life begun", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

var Chain *BlockChain = InitBlockChain()