package main

import "fmt"

type customErr struct {
	info string
}

func (e *customErr) Error() string {
	return e.info
}

func main() {
	err := customErr{
		"This is a custom error",
	}

	foo(&err)
}

func foo(e error) {
	fmt.Println(e)
}
