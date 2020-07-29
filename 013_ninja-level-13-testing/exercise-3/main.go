package main

import (
	"fmt"

	"github.com/toralfs/golang-basic-course/013_ninja-level-13-testing/exercise-3/mymath"
)

func main() {
	xxi := gen()

	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

// gen generates values to pass into CenteredAvg
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
