package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Transaction struct {
	Hash   string  `json:"hash"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Fee    float64 `json:"fee"`
	Sign   string  `json:"sign"`
}

func NewTransaction(privateKey *rsa.PrivateKey, from string, to string, amount float64) *Transaction {
	tx := Transaction{From: from, To: to, Amount: amount - (amount * 0.0001), Fee: amount * 0.0001}
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s %s %f %f", tx.From, tx.To, tx.Amount, tx.Fee)))
	tx.Hash = fmt.Sprintf("%x", hash)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}
	tx.Sign = base64.StdEncoding.EncodeToString(signature)
	return &tx
}

func (tx *Transaction) VerifyTransaction(pub *rsa.PublicKey, sign string, checksum string) error {
	signature, _ := base64.StdEncoding.DecodeString(sign)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, []byte(checksum), signature)
}
