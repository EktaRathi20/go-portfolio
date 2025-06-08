package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			println("Hello from goroutine", i)
		}(i)
	}

	// NOT BEST PRACTICE: Wait for goroutines to finish
	time.Sleep(2 * time.Second)

	// Best practice: Use a WaitGroup to wait for goroutines to finish
}
