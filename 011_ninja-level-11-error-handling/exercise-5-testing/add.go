package main

import "fmt"

func main() {
	fmt.Println(add(2, 4))
}

func add(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}
