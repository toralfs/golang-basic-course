package main

import (
	"fmt"
)

func main() {
	x := factorial(4)
	y := factorial2(4)
	fmt.Println(x)
	fmt.Println(y)
}

func factorial(n int) int { // From course
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

func factorial2(i int) int { // The one I came up with myself
	sum := i
	for j := i - 1; j > 0; j-- {
		sum *= j
	}
	return sum
}
