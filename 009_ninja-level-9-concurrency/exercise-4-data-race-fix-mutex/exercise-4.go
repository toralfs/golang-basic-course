package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	incrementer := 0
	const gs = 100
	wg.Add(100)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementer
			//runtime.Gosched() // Makes sense to remove this when using mutex
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value: ", incrementer)
}
