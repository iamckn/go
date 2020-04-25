package main

import (
	"fmt"
)

var x = int(1)
var y = string("foo")

func main() {
	if x == 1 {
		fmt.Println("x is 1")
	} else {
		fmt.Println("x is not 1")
	}

	switch y {
	case "foo":
		fmt.Println("y is foo")
	case "bar":
		fmt.Println("y is bar")
	default:
		fmt.Println("Default")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}
