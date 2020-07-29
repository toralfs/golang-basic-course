package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Current OS: ", runtime.GOOS)
	fmt.Println("Current architecure: ", runtime.GOARCH)
}
