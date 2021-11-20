package blockchain

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/PizzaNode/PizzaCoin/hash"
)

const version string = "00"

type Wallet struct {
	PublicKey  *rsa.PublicKey  `json:"public_key"`
	PrivateKey *rsa.PrivateKey `json:"private_key"`
}
type walletdto struct {
	PrivateKey string
	PublicKey  string
}

func LoadWallet() Wallet {
	_, err := os.Stat("wallet.json")
	if os.IsNotExist(err) {
		return NewWallet()
	}
	wallet_for := walletdto{}
	file, _ := os.ReadFile("wallet.json")
	json.Unmarshal(file, &wallet_for)
	privs, _ := base64.StdEncoding.DecodeString(wallet_for.PrivateKey)
	pubs, _ := base64.StdEncoding.DecodeString(wallet_for.PublicKey)
	privatekey, _ := x509.ParsePKCS1PrivateKey(privs)
	publickey, _ := x509.ParsePKCS1PublicKey(pubs)
	wallet := Wallet{
		publickey,
		privatekey,
	}
	return wallet
}

func NewWallet() Wallet {
	file, _ := os.Create("wallet.json")
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := privatekey.PublicKey
	wallet := Wallet{&publicKey, privatekey}

	wallet_for := walletdto{
		base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privatekey)),
		base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&publicKey)),
	}

	wallet_json, _ := json.Marshal(wallet_for)
	file.Write([]byte(wallet_json))
	file.Close()
	return wallet
}

func (w *Wallet) GetAddress() string {
	pub := x509.MarshalPKCS1PublicKey(w.PublicKey)
	vpayload := append([]byte(version), pub...)
	checksum := GetChecksum(vpayload)
	res := append(vpayload, checksum...)
	address := hash.GetSHA256(res)
	return address
}

func HashPublicKey(publicKey []byte) []byte {
	shasum := sha256.Sum256(publicKey)
	mdsum := md5.Sum(shasum[:])
	return mdsum[:]
}

func GetChecksum(in []byte) []byte {
	round1 := md5.Sum(in)
	round2 := md5.Sum(round1[:])
	return round2[:8]
}
