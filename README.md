# p2psch
A Peer-to-Peer network enabled chatroom for evading school IP blocking. The users connect to eachother instead of a centralized server so that school computers cannot block the IPs of said centralized serer.

## Disclaimer
I, the developer of p2psch, cannot be held accountable for the actions of the applications users as the application is merely a proof of concept. I do not condone or participate in breaking school rules. **USE AT YOUR WON RISK!**

## A note about Chromebooks and ChromeOS
If your school has given you Chromebooks and you want to use this application, you'd need to set ChromeOS in Developer Mode ([guide](https://www.howtogeek.com/210817/how-to-enable-developer-mode-on-your-chromebook/)) and install go manually ([guide](https://github.com/golang/go/wiki/ChromeOS)). One way to circumvent this would be to port the entire thing to JS, but I'm not sure pure JS's networking capabilities are great. There are ways to [get local IPs](https://tools.ietf.org/html/draft-ietf-rtcweb-security-arch-07#section-5.4) ([proof of concept](net.ipcalf.com)), but I cannot say for sure if it will work. Public IPs should also be pretty simple and can probably be implemented in the same fashion this project does. If someone wants to do this in JS, feel free to fork!

## Usage

## Internals
In the following subsections, I will describe the internals of p2psch.

### Networking
The network uses an inconsistent Peer-to-Peer model, meaning the network doesn't attempt to keep its state forever, much like [Snapchat](https://www.snapchat.com/).  The network's topology is a fully connected one (everyone is connected to everyone). Each student/device/**node** (these three terms will be used interchangeably) has the following parameters:
```golang
// EXCERPT FROM networking/node.go
type Node struct {
	ip string // A string representation of the node's current IP
	name string // The actual name the node shows in the chatroom
	messages []Message // The messages owned by the user
	contacts []Node // The peers/contacts the node is aware of
}
```