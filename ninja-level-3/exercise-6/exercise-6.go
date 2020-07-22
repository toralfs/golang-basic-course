package main

import "fmt"

func main() {
	fmt.Println("Even numbers between 0 and 100: ")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
