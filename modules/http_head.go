package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("https://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}

	resp.Body.Close()
	fmt.Println(resp.Status)
}
