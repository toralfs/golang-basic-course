package main

import "fmt"

func main() {
	dealer := struct {
		name      string
		inventory map[string]int
		employees []string
	}{
		name: "Scrooge's amazing used car dealership",
		inventory: map[string]int{
			"Ford":   10,
			"Toyota": 3,
			"Nissan": 5,
		},
		employees: []string{"Donald", "Dolly", "Scrooge"},
	}

	fmt.Printf("Welcome to %v \n\n", dealer.name)
	for k, v := range dealer.inventory {
		fmt.Printf("Inventory of %v: %v \n", k, v)
	}

	for i, v := range dealer.employees {
		fmt.Printf("Employee %v: %v \n", i, v)
	}
}
