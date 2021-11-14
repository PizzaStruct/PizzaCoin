package blockchain

type Block struct {
	ID          uint64
	Transaction []Transaction
	Hash        string
	PrevHash    string
	Timestamp   int64
	Miner       string
	Reward      uint64
	Fee         uint64
	Size        int64
}
