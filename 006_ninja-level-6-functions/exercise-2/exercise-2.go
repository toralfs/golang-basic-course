package main

import "fmt"

func foo(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(ii...)
	y := bar(ii)
	fmt.Println(x)
	fmt.Println(y)
}
