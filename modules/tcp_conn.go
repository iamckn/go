package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}

	_, err2 := net.Dial("tcp", "192.168.100.1:80")
	if err2 == nil {
		fmt.Println("Connection successful")
	}
}
