package networking

import ("github.com/r3tr0a11ways/chat/message.go"
		"github.com/r3tr0a11ways/networking/utils.go"
		"crypto/sha256")

const INITIAL_MESSAGE_SPACE = 64
const INITIAL_CONTACT_SPACE = 16
		
		
type Node struct {
	id []byte // The ID is created by sha256(users IP)
	ip string // A string representation of the node's current IP
	name string // The actual name the node shows in the chatroom
	messages []Message // The messages owned by the user
	contacts []Node // The peers/contacts the node is aware of
}

func (n *Node) Init(name string) {
	n.ip = utils.GetIP()
	n.id = sha256.Sum256([]byte(n.ip))
	n.name = name
	n.messages = make([]Message, INITIAL_MESSAGE_SPACE)
	n.contacts = make([]Node, INITIAL_CONTACT_SPACE)
}