package main

import (
	"fmt"
	"time"
)

// sendOnly uses a write-only channel parameter.
func sendOnly(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

func main() {
	// Buffered channel allows sender to proceed without immediate receiver.
	ch := make(chan int, 2)
	go sendOnly(ch)

	for v := range ch {
		fmt.Println("recv:", v)
	}

	timeout := time.After(50 * time.Millisecond)
	// select is the core primitive for multiplexing channel operations.
	select {
	case v := <-ch:
		fmt.Println("unexpected value:", v)
	case <-timeout:
		fmt.Println("select timeout")
	}
}
