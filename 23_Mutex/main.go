package main

import (
	"fmt"
	"sync"
)

// mutex is used to prevent race condition

type post struct {
	views int
	mu    sync.Mutex // Mutex to protect the views field
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock() // Unlock the mutex after updating the views field
	}()

	p.mu.Lock() // Lock the mutex before accessing the views field
	p.views++
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)                     // Increment the WaitGroup counter
		go myPost.incrementViews(&wg) // Increment views concurrently
	}

	wg.Wait()
	fmt.Println("Post views:", myPost.views)
}
