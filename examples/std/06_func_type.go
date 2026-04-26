package main

import "fmt"

// calculator is a named function type.
type calculator func(int, int) int

// apply accepts behavior as argument (higher-order function).
func apply(a, b int, fn calculator) int {
	return fn(a, b)
}

// choose returns different functions by operation name.
func choose(op string) calculator {
	switch op {
	case "add":
		return func(a, b int) int { return a + b }
	case "mul":
		return func(a, b int) int { return a * b }
	default:
		return func(a, b int) int { return 0 }
	}
}

func main() {
	add := func(a, b int) int { return a + b }
	fmt.Println("apply add:", apply(2, 3, add))
	fmt.Println("choose add:", apply(4, 5, choose("add")))
	fmt.Println("choose mul:", apply(4, 5, choose("mul")))
}
