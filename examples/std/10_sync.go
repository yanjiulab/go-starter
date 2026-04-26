package main

import (
	"fmt"
	"sync"
)

func main() {
	// WaitGroup waits all goroutines finish.
	var wg sync.WaitGroup
	// Mutex protects shared variable.
	var mu sync.Mutex
	// Once guarantees one-time initialization.
	var once sync.Once
	counter := 0
	initValue := 0

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				initValue = 100
			})

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter, "initValue:", initValue)
}
