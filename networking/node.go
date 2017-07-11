package networking

import ("github.com/r3tr0a11ways/p2psch/chat")

const INITIAL_MESSAGE_SPACE = 64
const INITIAL_CONTACT_SPACE = 16
		
		
type Node struct {
	Ip string // A string representation of the node's current IP
	Name string // The actual name the node shows in the chatroom
	Messages []chat.Message // The messages owned by the user
	Contacts []Node // The peers/contacts the node is aware of
}

func (n *Node) Init(name string) {
	n.Ip = GetLocalIP()
	n.Name = name
	n.Messages = make([]chat.Message, INITIAL_MESSAGE_SPACE)
	n.Contacts = make([]Node, INITIAL_CONTACT_SPACE)
}