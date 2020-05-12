package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Serialize(block *Block) []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(block)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

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