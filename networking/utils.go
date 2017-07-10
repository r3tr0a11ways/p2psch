package utils

import ("net/http"
		"bytes")
		
func GetIP() string{
	resp, _ := http.Get("http://ipv4.myexternalip.com/raw")
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return but.String()
}