package main

import ("github.com/r3tr0a11ways/p2psch/networking"
		"fmt")

func main() {
	var n *networking.Node = new(networking.Node)
	n.Init("Test")
	fmt.Printf("%v, %v, %v, %v\n", n.Ip, n.Name, n.Messages, n.Contacts)
}