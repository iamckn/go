package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)
}
