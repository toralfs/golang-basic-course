package main

import (
	"fmt"
	"sort"
)

// User defines a user
type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// SortByAge sorts a slice of users by their age
type SortByAge []User

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// SortByLast sorts a slice of users by their last name
type SortByLast []User

func (l SortByLast) Len() int           { return len(l) }
func (l SortByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l SortByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println("---------- Unsorted -------------")
	for _, u := range users {
		fmt.Printf("%s %s %d\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Printf("\t %s\n", v)
		}
		fmt.Println()
	}

	fmt.Println("---------- Sorted by age -------------")

	sort.Sort(SortByAge(users))
	for _, u := range users {
		fmt.Printf("%s %s %d\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Printf("\t %s\n", v)
		}
		fmt.Println()
	}

	fmt.Println("---------- Sorted by last name -------------")

	sort.Sort(SortByLast(users))
	for _, u := range users {
		fmt.Printf("%s %s %d\n", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Printf("\t %s\n", v)
		}
		fmt.Println()
	}
}
