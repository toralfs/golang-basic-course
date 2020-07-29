package main

import (
	"fmt"

	"github.com/toralfs/golang-basic-course/013_ninja-level-13-testing/exercise-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
