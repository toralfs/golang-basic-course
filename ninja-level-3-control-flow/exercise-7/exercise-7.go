package main

import "fmt"

func main() {
	even := make([]int, 0, 50)
	odd := make([]int, 0, 50)
	other := make([]int, 0, 5)
	for i := -5; i <= 100; i++ {
		if i > 0 {
			if i%2 == 0 {
				even = append(even, i)
			} else if i%2 != 0 {
				odd = append(odd, i)
			}
		} else {
			other = append(other, i)
		}
	}
	fmt.Println("Positive even numbers: ", even)
	fmt.Println("Positive odd numbers: ", odd)
	fmt.Println("Negative numbers: ", other)
}
