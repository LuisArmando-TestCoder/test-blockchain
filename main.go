package main

import (
	"github.com/luisarmando-testcoder/test-blockchain/blockchain"
	"github.com/luisarmando-testcoder/test-blockchain/server"
)

func main() {
	blockchain.Chain.AddBlock("You are kidding")
	server.InitServer()
}
