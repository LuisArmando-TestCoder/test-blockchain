package server

import (
	"encoding/json"
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
	encodeResponseAsJSON(blockchain.Chain.Blocks, w)
}

func GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blockHash := params["hash"]
	for _, block := range blockchain.Chain.Blocks {
		if (block.Hash == blockHash) {
			encodeResponseAsJSON(block, w)
			return
		}
	}
}

func PostBlock(w http.ResponseWriter, r *http.Request) {
	var data Block
	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &data)
	newBlock := blockchain.Chain.AddBlock(data.Data)

	encodeResponseAsJSON(newBlock, w)
}

func InitServer() {
	router := mux.NewRouter()

	router.HandleFunc("/blockchain", GetBlockchain).Methods("GET")
	router.HandleFunc("/block/{hash}", GetBlockByHash).Methods("GET")
	router.HandleFunc("/block", PostBlock).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}