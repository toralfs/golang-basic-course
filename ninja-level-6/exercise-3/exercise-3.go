package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("But this prints even later")
	}()
	fmt.Println("This prints last")
}

func bar() {
	fmt.Println("This prints first")
}

func main() {
	defer foo()
	bar()
}
