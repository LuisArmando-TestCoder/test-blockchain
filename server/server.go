package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/luisarmando-testcoder/test-blockchain/blockchain"
)

type Block struct {
	Data string  `json:"data,omitempty"`
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	json.NewEncoder(w).Encode(data)
}

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	DBBlockchain := blockchain.RetrieveBlockchain()
	encodeResponseAsJSON(DBBlockchain, w)
}

func GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blockHash := params["hash"]
	block := blockchain.RetrieveBlock(blockHash)
	encodeResponseAsJSON(block, w)
}

func PostBlock(w http.ResponseWriter, r *http.Request) {
	var data Block
	body, err := ioutil.ReadAll(r.Body)

	blockchain.Handle(err)

	unmarshalErr := json.Unmarshal(body, &data)

	blockchain.Handle(unmarshalErr)

	newBlock := blockchain.Chain.AddBlock(data.Data)

	blockchain.InsertBlock(newBlock)

	encodeResponseAsJSON(newBlock, w)
}

func InitServer() {
	router := mux.NewRouter()

	router.HandleFunc("/blockchain", GetBlockchain).Methods("GET")
	router.HandleFunc("/block/{hash}", GetBlockByHash).Methods("GET")
	router.HandleFunc("/block", PostBlock).Methods("POST")

	blockchain.OpenDatabase()

	blockchain.Chain.AddBlock("Am I kidding?")
	blockchain.InsertStartingBlockchainIfNeeded()

	fmt.Println(blockchain.RetrieveBlockchain())

	log.Fatal(http.ListenAndServe(":3000", router))
}