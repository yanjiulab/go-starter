package main

import (
	"fmt"
	"time"
)

// worker simulates asynchronous work.
func worker(id int, done chan<- int) {
	time.Sleep(time.Duration(id*30) * time.Millisecond)
	fmt.Println("worker done:", id)
	done <- id
}

func main() {
	// Launch multiple goroutines and collect completion signals.
	done := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("received completion from:", <-done)
	}
}
