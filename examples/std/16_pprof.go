package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

// busyWork creates CPU pressure for pprof demo.
func busyWork() {
	for i := 0; i < 5_000_000; i++ {
		_ = i * i
	}
}

func main() {
	fmt.Println("pprof: http://localhost:6060/debug/pprof/")
	fmt.Println("try: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=5")

	// CPU load goroutine.
	go func() {
		for {
			busyWork()
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// Memory allocation goroutine.
	go func() {
		for {
			_ = make([]byte, 1<<20)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Runtime status monitor.
	go func() {
		for {
			fmt.Println("goroutines:", runtime.NumGoroutine())
			time.Sleep(2 * time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe("localhost:6060", nil))
}
