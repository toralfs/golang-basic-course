package main

import "fmt"

func foo() func() {
	return func() {
		fmt.Println("This is printed from a function returned from another function")
	}
}

func main() {
	x := foo()

	x()
}
