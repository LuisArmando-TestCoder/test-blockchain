package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) *Block {
	lastHash, err := DB.Get([]byte(lastHashStr), nil)
	prevBlock := Chain.Blocks[len(Chain.Blocks)-1]

	if err == nil {
		Handle(err)
		prevBlock = RetrieveBlock(string(lastHash))
	}
	
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
	return new
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

var Chain *Blockchain = InitBlockchain()