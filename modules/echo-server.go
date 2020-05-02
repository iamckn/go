package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	//Create a buffer to store received data
	b := make([]byte, 512)
	for {
		//Receive data via conn
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		//Send data via conn
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}
func main() {
	//Bind to tcp port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		//Wait for connection to be established
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		//Use goroutine for concurrency
		go echo(conn)
	}
}
