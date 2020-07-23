package main

import "fmt"

type person struct {
	first  string
	last   string
	favIce []string
}

func main() {
	p1 := person{
		first:  "Donald",
		last:   "Duck",
		favIce: []string{"vanilla", "chocolate"},
	}

	p2 := person{
		first:  "Dolly",
		last:   "Duck",
		favIce: []string{"strawberry", "blueberry"},
	}

	sp := []person{p1, p2}

	for _, v := range sp {
		fmt.Println("name: ", v.first, v.last)
		for _, v2 := range v.favIce {
			fmt.Printf("\t %v \n", v2)
		}
	}
}
