package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("This is printed by a func assigned to a variable")
	}

	x()
}
