package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10 // number of goroutines we want
	y := 10 // number of ints put on the channel by each goroutine
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	return c
}
