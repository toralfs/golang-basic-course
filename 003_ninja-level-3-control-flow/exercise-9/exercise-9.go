package main

import "fmt"

func main() {

	favSport := "parkour"

	switch favSport {
	case "snowboarding":
		fmt.Println("my favorite sport is done in the snow")
	case "hockey":
		fmt.Println("my favorite sport is done on the ice")
	case "E-GAMING":
		fmt.Println("my favorte sport is done on the computer TASTELESS")
	case "parkour":
		fmt.Println("my favorite sport is done where ever")
	default:
		fmt.Println("now my favorite sport is something different, something unknown to the minds of men")
	}
}
