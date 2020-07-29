package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["someone else"] = []string{"something else", "something fun", "something interesting"}
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("record for: ", k)
		for i, v2 := range v {
			fmt.Printf("\t index in slice: %v; value: %v\n", i, v2)
		}
	}
}
