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
		first:  "Scrooge",
		last:   "McDuck",
		favIce: []string{"pistachio", "blueberry"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, k := range m {
		fmt.Println("name: ", k.first, k.last)
		for _, v := range k.favIce {
			fmt.Printf("\t %v \n", v)
		}
	}
}
