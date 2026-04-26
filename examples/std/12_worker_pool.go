package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	worker int
	value  int
}

// worker consumes jobs and publishes results.
func worker(id int, jobs <-chan int, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(20 * time.Millisecond)
		results <- result{worker: id, value: j * j}
	}
}

func main() {
	// Queue jobs first, then close channel so workers can exit naturally.
	jobs := make(chan int, 5)
	results := make(chan result, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, job := range []int{1, 2, 3, 4, 5} {
		jobs <- job
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Printf("worker-%d result=%d\n", r.worker, r.value)
	}
}
