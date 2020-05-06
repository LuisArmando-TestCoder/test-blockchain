package blockchain

import "strconv"

type Block struct {
	Hash	 string
	Data	 string
	PrevHash string
}

type BlockChain struct {
	Blocks []*Block
}

func getDataHash(data string) int {
	var hash int
	for i := 0; i < len(data); i++ {
		character := int([]rune(string(data[i]))[0])
		hash = ((hash << 5) - hash) + character
		hash = hash & hash
	}
	return hash;
}

func (block *Block) InsertHash() {
	hash := getDataHash(block.Data + block.PrevHash)
	block.Hash = strconv.Itoa(hash)
}

func CreateBlock(data string, prevHash string) *Block {
	block := &Block{"", data, prevHash}
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
	return CreateBlock("The life begun", "with a click")
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

var Chain *BlockChain = InitBlockChain()