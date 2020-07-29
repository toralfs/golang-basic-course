package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type human interface {
	speak(s string)
}

func (p *person) speak(s string) {
	fmt.Printf("Person: %v %v says: %s \n", p.First, p.Last, s)
}

func saySomething(h human) {
	h.speak("Quack Quack")
}

func main() {
	p1 := person{
		First: "Dolly",
		Last:  "Duck",
		Age:   30,
	}

	p1.speak("Quack")
	saySomething(&p1) // Works
	//saySomething(p1)  // Does not work
	human.speak(&p1, "Quack Quack Quack") // Works
	//human.speak(p1, "Quack Quack Quack") // Does not work
}
