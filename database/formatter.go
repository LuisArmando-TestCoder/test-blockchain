package database

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/luisarmando-testcoder/test-blockchain/blockchain"
)

func Serialize(block *blockchain.Block) []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(block)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *blockchain.Block {
	var block blockchain.Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}