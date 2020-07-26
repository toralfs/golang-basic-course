package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 14, 15, 16, 17, 18} // composite literal

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", a)
}
