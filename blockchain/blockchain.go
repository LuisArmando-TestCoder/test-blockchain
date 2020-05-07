package blockchain

type Block struct {
	Hash	 string
	Data	 string
	PrevHash string
}

type BlockChain struct {
	Blocks []*Block
}

func (block *Block) InsertHash() {
	block.Hash = GetProvenHash(block)
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