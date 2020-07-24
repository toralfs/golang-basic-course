package main

import "fmt"

func sum(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func oddSum(f func(ii ...int) int, xi ...int) int {
	var oi []int
	for _, v := range xi {
		if v%2 != 0 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}

func evenSum(f func(ii ...int) int, xi ...int) int {
	var oi []int
	for _, v := range xi {
		if v%2 == 0 {
			oi = append(oi, v)
		}
	}
	return f(oi...)
}

func main() {
	x := oddSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	y := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(x)
	fmt.Println(y)
}
