package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person, newFirst string, newLast string, newAge int) {
	p.first = newFirst // p.first == (*p).first
	p.last = newLast
	p.age = newAge
}

func main() {
	p1 := person{
		first: "Dolly",
		last:  "Duck",
		age:   30,
	}

	fmt.Println(p1)
	changeMe(&p1, "Donald", "Duck", 32)
	fmt.Println(p1)
}
