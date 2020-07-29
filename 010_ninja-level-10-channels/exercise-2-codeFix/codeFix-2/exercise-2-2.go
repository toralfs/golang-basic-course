package main

import (
	"fmt"
)

// Base code that doesn't work
// func main() {
// 	cr := make(<-chan int)

// 	go func() {
// 		cr <- 42
// 	}()
// 	fmt.Println(<-cr)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cr\t%T\n", cr)
// }

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
