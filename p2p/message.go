package p2p

import "net"

// Messae holds any arbitrary data that is being sent over each transport between two nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
