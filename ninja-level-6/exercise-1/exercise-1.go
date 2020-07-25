package main

import "fmt"

func foo(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar() (string, int) {
	return "The answer to everything: ", 42
}

func main() {
	x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	t, n := bar()
	fmt.Println(x)
	fmt.Println(t, n)
}
