package main

import "fmt"

var fruits []string = []string{"apple", "orange", "pear", "banana"}

func isBanana(s string) bool {
	if s == "banana" {
		return true
	}
	return false
}

func bananaSearch(f func(s string) bool, ss []string) {
	for i, v := range ss {
		if f(v) {
			fmt.Printf("banana found at index %v\n", i)
			return
		}
	}
	fmt.Println("No bananas found :(")
}

func main() {
	bananaSearch(isBanana, fruits)
}
