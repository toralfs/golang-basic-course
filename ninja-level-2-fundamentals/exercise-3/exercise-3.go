package main

import "fmt"

const (
	a      = 1  // untyped
	b int8 = 10 // typed
)

func main() {
	fmt.Printf("type for a :%T\ttype for b: %T\n", a, b)
}
