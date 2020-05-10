package database

import (
	"github.com/luisarmando-testcoder/test-blockchain/blockchain"
	"github.com/dgraph-io/badger"
)

const LAST_HASH = "lastHash"
const dbPath = "./tmp/badger"

var DB *badger.DB

	// Dir = dbPath
	// where the db will store the keys and metadata
	// ValueDir = dbPath
	// where the db will store all of the values
func OpenDatabase() {

	opts := badger.DefaultOptions(dbPath).WithSyncWrites(true)

	db, err := badger.Open(opts)

	Handle(err)

	DB = db
}

func GetLastHash() string {

	var lastHash string

	err := DB.View(func (txn *badger.Txn) error {
		item, err := txn.Get([]byte(LAST_HASH))

		Handle(err)

		if err == nil {
			item.Value(func(val []byte) error {
				lastHash = Deserialize(val).Hash
				return nil
			})
		}

		return err
	})

	Handle(err)

	return lastHash
}


func getBlock(txn *badger.Txn, lastHash string) func() (item *badger.Item, rerr error) {
	return func() (item *badger.Item, rerr error) {
		return txn.Get([]byte(lastHash))
	}
}

func RetrieveBlockchain() []*blockchain.Block {
	var tempBlockchain []*blockchain.Block
	DB.View(func (txn *badger.Txn) error {

		lastHashItem, err := txn.Get([]byte(LAST_HASH))

		if err == nil {
			lastHashItem.Value(func(val []byte) error {
				lastHash := Deserialize(val).Hash

				gb := getBlock(txn, lastHash)

				for blockItem, err := gb(); err == nil; blockItem, err = gb() {
					blockItem.Value(func(val []byte) error {
						block := Deserialize(val)
						tempBlockchain = append(blockchain.Chain.Blocks, block)
						return nil
					})
				}
				return nil
			})
		}

		Handle(err)

		return err
	})
	
	return tempBlockchain
}

func RetrieveBlock(hash string) *blockchain.Block {
	var block *blockchain.Block
	DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(hash))
		if err == nil {
			item.Value(func(val []byte) error {
				block = Deserialize(val)
				return nil
			})
		}
		Handle(err)
		return err
	})
	return block
}

func InsertStartingBlockchainIfNeeded() {
	var err error 
	err = DB.Update(func(txn *badger.Txn) error {
		blocks := blockchain.Chain.Blocks
		for _, block := range blocks {
			_, err := txn.Get([]byte(block.Hash))

			Handle(err)

			if err != nil {
				txn.Set([]byte(block.Hash), Serialize(block))
			}
		}

		lastHashKey := []byte(LAST_HASH)
		lastHashValue := []byte(blocks[len(blocks)-1].Hash)

		err = txn.Set(lastHashKey, lastHashValue)

		return err
	})

	Handle(err)
}

func InsertBlock(block blockchain.Block) {
	var err error 
	err = DB.Update(func(txn *badger.Txn) error {
		txn.Set([]byte(block.Hash), Serialize(&block))
		lastHashKey := []byte("lastHash")
		lastHashValue := []byte(block.Hash)

		err = txn.Set(lastHashKey, lastHashValue)

		return err
	})

	Handle(err)
}