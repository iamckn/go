package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		remoteaddress := net.TCPAddr{
			Port: i,
			IP:   net.ParseIP("45.33.32.156"),
		}
		conn, err := net.DialTCP("tcp", nil, &remoteaddress)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
