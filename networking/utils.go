package networking

import ("net/http"
		"net"
		"bytes"
		"log"
		"strings")
		
func GetPublicIP() string {
	resp, err := http.Get("http://ipv4.myexternalip.com/raw")
	if err != nil {
		log.Panic("Cannot get user's Public IP")
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Panic(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().String()
    idx := strings.LastIndex(localAddr, ":")
    return localAddr[0:idx]
}