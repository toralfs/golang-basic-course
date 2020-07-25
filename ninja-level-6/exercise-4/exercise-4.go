package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %s %s, I am %d years old\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{"Dolly", "Duck", 30}
	p1.speak()
}
