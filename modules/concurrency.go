package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Ran first function")
}

func main() {
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("Ran main function")
}
