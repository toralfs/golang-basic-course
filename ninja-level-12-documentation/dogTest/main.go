package main

import (
	"fmt"

	"github.com/toralfs/golang-basic-course/ninja-level-12-documentation/dog"
)

type animal struct {
	name string
	age  int
}

func main() {
	frøya := animal{
		name: "Frøya",
		age:  dog.Years(6),
	}
	fmt.Println(frøya)
}
