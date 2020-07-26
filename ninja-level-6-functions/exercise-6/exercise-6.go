package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is printed by an anonymous function")
	}()
}
