package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PizzaNode/PizzaCoin/blockchain"
	"github.com/PizzaNode/PizzaCoin/peer"
)

func ShareBlockchain(blocks []blockchain.Block) {
	for _, to := range peer.Peers {
		json_blocks, _ := json.Marshal(blocks)
		resp, err := http.Post(fmt.Sprintf("%s/block", to.Host), "application/json", bytes.NewBuffer(json_blocks))
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}

func SharePeer(per peer.Peer) {
	for _, to := range peer.Peers {
		json_peer, _ := json.Marshal(per)
		resp, err := http.Post(fmt.Sprintf("%s/peers", to.Host), "application/json", bytes.NewBuffer(json_peer))
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}

func ShareTx(tx blockchain.Transaction) {
	for _, to := range peer.Peers {
		json_tx, _ := json.Marshal(tx)
		resp, err := http.Post(fmt.Sprintf("%s/tx", to.Host), "application/json", bytes.NewBuffer(json_tx))
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
}
