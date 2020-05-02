package main

import (
	"fmt"
	"net"
	"sort"
)

//Function to do the scan using channels
func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("192.168.100.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	//Create the ports channel to receive 1000 values at a time
	ports := make(chan int, 1000)

	//Create the results channel
	results := make(chan int)

	//reate the slice to store results for later sorting
	var openports []int

	//Start the workers for scanning
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	//Start a routine to feed the ports channel
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	//Append open ports to the openports channel
	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
