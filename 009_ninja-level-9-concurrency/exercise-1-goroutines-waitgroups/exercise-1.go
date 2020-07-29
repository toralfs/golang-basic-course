package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("First go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Second go routine")
		wg.Done()
	}()

	fmt.Println("mid goroutines: ", runtime.NumGoroutine())
	fmt.Println("mid threads: ", runtime.NumCPU())

	wg.Wait()

	fmt.Println("end goroutines: ", runtime.NumGoroutine())
	fmt.Println("end threads: ", runtime.NumCPU())
}
