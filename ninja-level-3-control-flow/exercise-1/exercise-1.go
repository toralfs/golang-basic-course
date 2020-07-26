// Print every number from 1 to 10 000

package main

import "fmt"

func main() {
	n := make([]int, 0, 10000)

	for i := 1; i <= cap(n); i++ {
		n = append(n, i)
	}

	fmt.Println(n)
}

/*
func main() {
	n := make([]int, 0, 10000)

	for i := 1; i <= cap(n); i++ {
		n = append(n, i)
	}

	for j := range n {
		fmt.Printf("%d ", j)
		if j%10 == 0 {
			fmt.Printf("\n")
		}
	}
}
*/
