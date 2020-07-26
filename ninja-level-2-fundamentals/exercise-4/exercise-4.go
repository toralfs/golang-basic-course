package main

import "fmt"

var x int

func main() {
	x = 1
	fmt.Printf("decimal: %d\tbinary: %b\thex: %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("decimal: %d\tbinary: %b\thex: %#x\n", y, y, y)
}
