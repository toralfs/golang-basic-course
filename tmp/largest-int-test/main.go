package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func largest(ii []int) int {
	var large int
	for _, i := range ii {
		if i > large {
			large = i
		}
	}
	return large
}

func largest2(ii []int) int {
	sort.Ints(ii)
	return ii[len(ii)-1]
}

func randNum() []int {
	var ii []int
	for i := 0; i < 1000; i++ {
		rand.Seed(2)
		num := rand.Intn(1000)
		ii = append(ii, num)
	}
	return ii
}

func main() {
	ii := randNum()
	fmt.Println(largest(ii))
	fmt.Println(largest2(ii))
}
