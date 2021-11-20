package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/PizzaNode/PizzaCoin/blockchain"
	"github.com/PizzaNode/PizzaCoin/helpers"
	"github.com/PizzaNode/PizzaCoin/peer"
	"github.com/PizzaNode/PizzaCoin/peer/client"
	"github.com/joho/godotenv"
)

var (
	Bc      = blockchain.LoadBlockchain()
	Wallet  = blockchain.LoadWallet()
	Address = Wallet.GetAddress()
)

func Run() {
	if err := godotenv.Load("./config/.env"); err != nil {
		panic(err)
	}
	_, ok := os.LookupEnv("PIZZACOIN_ROOT")
	if !ok {
		fmt.Fprintf(os.Stderr, "Please setup a PIZZACOIN_ROOT variable\n")
		os.Exit(1)
	}

	args := os.Args[1:]
	switch args[0] {
	case "ping":
		fmt.Println("pong")
	case "mine":
		helpers.Loading(context.TODO(), "Mining in process... Type ^C to stop mine")
		for {
			Bc.NewBlock(Address)
			client.ShareBlockchain(Bc.Blocks)
			helpers.UpdateJsonFile(Bc, "blockchain.json")
		}
	case "newtx":
		amount, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			panic(err)
		}
		tx := Bc.NewTransaction(Wallet.PrivateKey, Address, args[1], amount)
		client.ShareTx(tx)
	case "wallet":
		fmt.Printf("Wallet Address: %s\nBalance: %f PizzaCoins", Address, Bc.GetBalance(Address))
	case "blockchain":
		res, _ := json.MarshalIndent(Bc, "", "  ")
		fmt.Println(string(res))
	case "peers":
		for i := 0; i < len(peer.Peers); i++ {
			fmt.Printf("%d. %s", i+1, peer.Peers[i].Host.String())
		}
	case "connect":
		break
	case "disconnect":
		break
	default:
		fmt.Printf("Unknown argument - %s", os.Args[1])
		os.Exit(1)
	}
	os.Exit(0)
}
