package main

import "fmt"

type kiwis int

var x kiwis
var y int

func main() {
	fmt.Printf("value of x: %v; type of x: %T\n", x, x)
	x = 42
	fmt.Printf("value of x: %v\n", x)

	y = int(x)
	fmt.Printf("value of y: %v; type of y: %T\n", y, y)
}
