package p2p

// Peer is an interface that represents the remote node
type Peer interface{}

// Transport is anything that hundles communication between the nodes int the network
// can be of the form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
