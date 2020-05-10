package blockchain

type Block struct {
	Hash	 string
	Data	 string
	PrevHash string
}

func (block *Block) InsertHash() {
	block.Hash = GetProvenHash(block)
}

func CreateBlock(data string, prevHash string) *Block {
	block := &Block{"", data, prevHash}
	block.InsertHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("The life begun", "")
}