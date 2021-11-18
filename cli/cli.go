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
	"github.com/joho/godotenv"
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

	bc := blockchain.LoadBlockchain()
	wallet := peer.LoadWallet()
	address := wallet.GetAddress()

	args := os.Args[1:]
	switch args[0] {
	case "ping":
		fmt.Println("pong")
	case "mine":
		helpers.Loading(context.TODO(), "Mining in process... Type ^C to stop mine")
		for {
			bc.AddBlock(address)
			helpers.UpdateJsonFile(bc, "blockchain.json")
		}
	case "newtx":
		amount, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			panic(err)
		}
		bc.AddTransaction(wallet.PrivateKey, address, args[1], amount)
	case "wallet":
		fmt.Printf("Wallet Address: %s\nBalance: %f PizzaCoins", address, bc.GetBalance(address))
	case "blockchain":
		res, _ := json.MarshalIndent(bc, "", "  ")
		fmt.Println(string(res))
	case "peers":
		break
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
