package blockchain

import (
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
)


const dbPath = "./tmp/level"
const lastHashStr = "lastHash"

var DB *leveldb.DB

func OpenDatabase() {
	db, err := leveldb.OpenFile(dbPath, nil)

	Handle(err)

	DB = db
}

func RetrieveBlockchain() []Block {
	var tempBlockchain []Block

	iter := DB.NewIterator(nil, nil)

	for iter.Next() {
		_, err := strconv.Atoi(string(iter.Value()))
		if err != nil {
			value := *Deserialize(iter.Value())
			tempBlockchain = append(tempBlockchain, value)
		}
		Handle(err)
	}
	iter.Release()
	err := iter.Error()

	Handle(err)

	return tempBlockchain
}

func RetrieveBlock(hash string) *Block {
	block, err := DB.Get([]byte(hash), nil)

	Handle(err)

	return Deserialize(block)
}

func InsertBlock(block *Block) {
	err := DB.Put([]byte(block.Hash), Serialize(block), nil)
	DB.Put([]byte(lastHashStr), []byte(block.Hash), nil)
	Handle(err)
}

func InsertStartingBlockchainIfNeeded() {
	iter := DB.NewIterator(nil, nil)

	canUseDB := iter.Next()

	if canUseDB == false {
		for _, block := range Chain.Blocks {
			DB.Put([]byte(block.Hash), Serialize(block), nil)
		}
		lastBlock := Chain.Blocks[len(Chain.Blocks)-1]
		DB.Put([]byte(lastHashStr), []byte(lastBlock.Hash), nil)
	}
}