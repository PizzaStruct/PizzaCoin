package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PizzaNode/PizzaCoin/blockchain"
	"github.com/PizzaNode/PizzaCoin/cli"
	"github.com/PizzaNode/PizzaCoin/hash"
	"github.com/PizzaNode/PizzaCoin/peer"
	"github.com/PizzaNode/PizzaCoin/peer/client"

	"github.com/gorilla/mux"
)

func RunServer() {
	mux := GetRouter()
	server := http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func GetRouter() http.Handler {
	router := mux.NewRouter()
	bex_handler := BEXHandler{}
	pex_handler := PEXHandler{}
	router.HandleFunc("/tx", bex_handler.AddTransaction).Methods("POST")
	router.HandleFunc("/block", bex_handler.ReplaceBlockchain).Methods("POST")
	router.HandleFunc("/peers", pex_handler.AddPeer).Methods("POST")
	return router
}

type BEXHandler struct {
}

func (bex *BEXHandler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": 500}"))
		return
	}
	var tx blockchain.Transaction
	json.Unmarshal(body, &tx)
	cli.Bc.AddTransaction(&tx)
	client.ShareTx(tx)
	w.WriteHeader(202)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": 202}"))
}

func (bex *BEXHandler) ReplaceBlockchain(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": 500}"))
		return
	}
	var blocks []blockchain.Block
	json.Unmarshal(body, &blocks)
	cli.Bc.ReplaceBlockchain(blocks)
	client.ShareBlockchain(blocks)
	w.WriteHeader(202)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": 202}"))
}

type PEXHandler struct {
}

func (pex *PEXHandler) AddPeer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": 500}"))
		return
	}
	var per peer.Peer
	json.Unmarshal(body, &per)
	var found bool = false
	for _, v := range peer.Peers {
		if v.PeerHash == hash.GetSHA256([]byte(per.Host.String())) {
			found = true
			break
		}
	}
	if !found {
		peer.Peers = append(peer.Peers, per)
	}
	w.WriteHeader(202)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": 202}"))
}
