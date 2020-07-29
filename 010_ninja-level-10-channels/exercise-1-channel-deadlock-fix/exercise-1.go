package main

import (
	"fmt"
)

func main() {
	funcLiteral()
	bufferedChannel()
}

// Base code that is in dead lock
// func main() {
// 	c := make(chan int)

// 	c <- 42

// 	fmt.Println(<-c)
// }

func bufferedChannel() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func funcLiteral() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
