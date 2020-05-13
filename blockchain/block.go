package blockchain

type Block struct {
	Hash	 string
	Data	 string
	PrevHash string
	Nonce    int
}

func (block *Block) InsertHash() {
	hash, nonce := GetProvenHash(block)
	block.Hash = hash
	block.Nonce = nonce
}

func CreateBlock(data string, prevHash string) *Block {
	block := &Block{"", data, prevHash, 0}
	block.InsertHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("The life begun", "")
}