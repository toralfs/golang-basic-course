package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("value for x: %v \nvalue for y: %v \nvalue for z: %v \n", x, y, z)

	fmt.Println(s)
}
