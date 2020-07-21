package main

import "fmt"

const (
	_  = iota
	y1 = (2020 - 4) + iota
	y2 = (2020 - 4) + iota
	y3 = (2020 - 4) + iota
	y4 = (2020 - 4) + iota
)

func main() {
	fmt.Println(y1, y2, y3, y4)
}
