package message

import ("time")

type Message struct {
	timestamp Time // The time at which the message was recieved
	content string // The content of the message
	sender []byte // The ID of the sender
}