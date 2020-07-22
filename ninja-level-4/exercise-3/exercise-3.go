package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 14, 15, 16, 17, 18} // composite literal

	fmt.Println(a[:5])
	fmt.Println(a[5:])
	fmt.Println(a[2:7])
	fmt.Println(a[3:8])
}
