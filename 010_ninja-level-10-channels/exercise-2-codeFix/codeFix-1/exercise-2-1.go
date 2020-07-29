package main

import (
	"fmt"
)

// Base code that doesn't work
// func main() {
// 	cs := make(chan<- int)

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
