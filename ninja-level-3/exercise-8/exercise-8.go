package main

import "fmt"

func main() {

	i := 1293821938413

	switch {
	case i%2 == 0:
		fmt.Println("i is an even number: ", i)
	case i%2 != 0:
		fmt.Println("i is not an even number: ", i)
	}

	/* More simplified example
	switch {
	case true:
		fmt.Println("Should print")
	case false:
		fmt.Println("Should not print")
	}
	*/
}
