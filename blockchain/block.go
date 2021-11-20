package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PizzaNode/PizzaCoin/hash"
)

type Block struct {
	ID           uint64        `json:"id"`
	Transactions []Transaction `json:"transactions"`
	Hash         string        `json:"hash"`
	PrevHash     string        `json:"prev_hash"`
	Timestamp    int64         `json:"timestamp"`
	Miner        string        `json:"miner"`
	Reward       float64       `json:"reward"`
	Fee          float64       `json:"fee"`
	Nonce        uint64        `json:"nonce"`
}

func (b *Block) GetHash() string {
	txs, _ := json.Marshal(b.Transactions)
	block_plain := fmt.Sprintf(
		"%d-%s-%d-%s-%d-%s",
		b.ID,
		b.PrevHash,
		b.Nonce,
		b.Miner,
		time.Now().Unix(),
		txs,
	)
	return hash.GetSHA256([]byte(block_plain))
}

func NewBlock(id uint64, txs *[]Transaction, ph string, r float64, miner string) *Block {
	block := Block{
		ID:           id,
		Transactions: *txs,
		PrevHash:     ph,
		Reward:       r,
		Fee:          CollectFee(txs),
		Miner:        miner,
		Nonce:        0,
		Timestamp:    time.Now().Unix(),
	}
	block.Mine(6, txs)
	return &block
}

func CollectFee(txs *[]Transaction) float64 {
	var total float64 = 0
	for _, v := range *txs {
		total = v.Fee
	}
	return total
}

func (b *Block) Mine(d int, txs *[]Transaction) {
	diff := strings.Repeat("0", d)
	for !strings.HasPrefix(b.Hash, diff) {
		b.Transactions = *txs
		r_tx := Transaction{From: "", To: b.Miner, Amount: b.Reward + CollectFee(txs), Fee: 0}
		r_tx.Hash = string(hash.GetSHA256([]byte(fmt.Sprintf("%s %s %f %f", r_tx.From, r_tx.To, r_tx.Amount, r_tx.Fee))))
		b.Transactions = append(b.Transactions, r_tx)
		b.Nonce++
		b.Hash = b.GetHash()
	}
	fmt.Printf("\n\rSuccessfully mined block with hash %s nonce %d\n\r", b.Hash, b.Nonce)
}
