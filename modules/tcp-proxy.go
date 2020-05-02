package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	//Destination connection
	dst, err := net.Dial("tcp", "10.10.10.1:80")
	if err != nil {
		log.Fatalln("Unable to connect")
	}

	defer dst.Close()

	go func() {
		//copy source output to destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	//copy destination output back to source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//Listen on local port 80
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
