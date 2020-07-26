package main

import "fmt"

func main() {
	ss := [][]string{{"James", "Bond", "not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James"}}

	for i, r := range ss {
		fmt.Println("record: ", i)
		for j, d := range r {
			fmt.Printf("\t index: %v; value: %v\n", j, d)
		}
	}
}
