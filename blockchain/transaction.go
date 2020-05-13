package blockchain

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	ID 		int
	Inputs 	[]TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  int
	Out int
	Sig string
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{-1, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{-1, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) SetID() {
	out, err := json.Marshal(tx)

	Handle(err)

	tx.ID = getDataHash(string(out))
}