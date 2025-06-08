package main

import "sync"

func printGoroutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	println("Hello from goroutine", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go printGoroutine(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
}
