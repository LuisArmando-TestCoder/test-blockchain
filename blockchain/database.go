package blockchain

import "github.com/syndtr/goleveldb/leveldb"


const dbPath = "./tmp/level"

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
		value := Deserialize(iter.Value())
		
		tempBlockchain = append(tempBlockchain, *value)
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

	Handle(err)
}

func InsertStartingBlockchainIfNeeded() {
	iter := DB.NewIterator(nil, nil)

	canUseDB := iter.Next()

	if canUseDB == false {
		for _, block := range Chain.Blocks {
			DB.Put([]byte(block.Hash), Serialize(block), nil)
		}
	}
}