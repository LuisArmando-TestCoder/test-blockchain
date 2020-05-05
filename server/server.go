package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strconv"

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

func GetBlock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blockIndex, stringConversionError := strconv.Atoi(params["id"])
	block := blockchain.Chain.Blocks[blockIndex % len(blockchain.Chain.Blocks)]
	if block != nil && stringConversionError == nil {
		encodeResponseAsJSON(block, w)
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
	router.HandleFunc("/block/{id}", GetBlock).Methods("GET")
	router.HandleFunc("/block", PostBlock).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}