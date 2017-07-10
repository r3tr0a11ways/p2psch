# p2psch
Peer-to-Peer network that contains a small-scale social networking feature for it to be hosted on school computers. 

## Networking
The network uses a Peer-to-Peer model inspired by (IPFS)[ipfs.io], although with a significant difference - the network doesn't attempt to keep its state forever. The network is perfectly fine with losing all the data on it. Each student/device/*node* (these three terms will be used interchangeably) has the following parameters:
```
// EXCERPT FROM networking/node.go
type Node struct {
	id []byte // The ID is created by sha256(users IP)
	ip string // A string representation of the node's current IP
	name string // The actual name the node shows in the chatroom
	messages []Message // The messages owned by the user
	contacts []Node // The peers/contacts the node is aware of
}
```