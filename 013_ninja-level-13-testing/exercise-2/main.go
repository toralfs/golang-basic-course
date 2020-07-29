package main

import (
	"fmt"

	"github.com/toralfs/golang-basic-course/013_ninja-level-13-testing/exercise-2/quote"
	"github.com/toralfs/golang-basic-course/013_ninja-level-13-testing/exercise-2/word"
)

func main() {

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
