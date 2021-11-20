package blockchain

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"os"

	"github.com/PizzaNode/PizzaCoin/helpers"
)

type Blockchain struct {
	Blocks      []Block       `json:"blocks"`
	Pending_txs []Transaction `json:"pending_txs"`
}

func InitBlockchain() *Blockchain {
	bc := Blockchain{}
	return &bc
}

func LoadBlockchain() *Blockchain {
	file, err := os.ReadFile("blockchain.json")
	var bc Blockchain
	if err != nil {
		os.Create("blockchain.json")
		bc = *InitBlockchain()
		helpers.UpdateJsonFile(bc, "blockchain.json")
		return &bc
	}
	json.Unmarshal(file, &bc)
	return &bc
}

func (bc *Blockchain) NewBlock(miner string) {
	var reward float64 = 50
	if len(bc.Blocks) == 0 {
		b := NewBlock(0, &[]Transaction{}, "", reward, miner)
		bc.Blocks = append(bc.Blocks, *b)
		return
	}
	prev := bc.Blocks[len(bc.Blocks)-1]
	b := NewBlock(prev.ID+1, &bc.Pending_txs, prev.Hash, reward, miner)
	bc.Pending_txs = []Transaction{}
	bc.Blocks = append(bc.Blocks, *b)
	helpers.UpdateJsonFile(bc, "blockchain.json")
}

func (bc *Blockchain) ReplaceBlockchain(blocks []Block) {
	bc.Blocks = blocks
	helpers.UpdateJsonFile(bc, "blockchain.json")
}

func (bc *Blockchain) NewTransaction(pk *rsa.PrivateKey, from string, to string, amount float64) Transaction {
	tx := NewTransaction(pk, from, to, amount)
	bc.Pending_txs = append(bc.Pending_txs, *tx)
	helpers.UpdateJsonFile(bc, "blockchain.json")
	return *tx
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	bc.Pending_txs = append(bc.Pending_txs, *tx)
	helpers.UpdateJsonFile(bc, "blockchain.json")
}

func (bc *Blockchain) ValidateChain() error {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PrevHash != bc.Blocks[i-1].GetHash() {
			return errors.New("Blockchain is wrong")
		}
	}
	return nil
}

func (bc *Blockchain) GetBalance(address string) float64 {
	var res float64
	for i := 0; i < len(bc.Blocks); i++ {
		for j := 0; j < len(bc.Blocks[i].Transactions); j++ {
			if bc.Blocks[i].Transactions[j].From == address {
				res -= bc.Blocks[i].Transactions[j].Amount
			}
			if bc.Blocks[i].Transactions[j].To == address {
				res += bc.Blocks[i].Transactions[j].Amount
			}
		}
	}
	return res
}
