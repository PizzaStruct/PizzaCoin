package peer

import (
	"net"

	"github.com/PizzaNode/PizzaCoin/hash"
)

type Peer struct {
	PeerHash string
	Host     net.IP
}

func NewPeer(host net.IP) Peer {
	hash := hash.GetSHA256([]byte(host.String()))
	return Peer{Host: host, PeerHash: hash}
}

var Peers []Peer
