package chat

import ("time")

type Message struct {
	Timestamp time.Time // The time at which the message was recieved
	Content string // The content of the message
	Sender []byte // The ID of the sender
}