package blockchain

type Transaction struct {
	From   string
	To     string
	Amount uint64
	Fee    uint64
}
