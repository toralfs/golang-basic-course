package main

import "fmt"

type kiwis int

var x kiwis

func main() {
	fmt.Printf("value of x: %v; type of x: %T\n", x, x)

	x = 42

	fmt.Printf("value of x: %v;", x)
}
