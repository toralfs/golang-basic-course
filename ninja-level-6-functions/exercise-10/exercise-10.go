// Closure, incrementer example

package main

import "fmt"

func foo() func() int {
	x := 0 // encloses x
	return func() int {
		x++
		return x
	}
}

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
