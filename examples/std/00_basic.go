package main

import "fmt"

// add shows a simple function with typed parameters.
func add(a, b int) int {
	return a + b
}

// divide demonstrates multiple return values and error handling.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Constants and short variable declarations.
	const pi = 3.14
	name := "gopher"
	age := 3
	score := 98.5

	// Branching with if/else.
	if age >= 3 {
		fmt.Println(name, "is growing with Go")
	} else {
		fmt.Println(name, "is too young")
	}

	// for is Go's only loop keyword.
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += i
	}

	switch name {
	case "gopher":
		fmt.Println("official mascot")
	default:
		fmt.Println("unknown")
	}

	defer fmt.Println("defer runs at function end")
	fmt.Println("pi:", pi, "score:", score, "sum:", sum, "add:", add(2, 5))

	if q, err := divide(10, 2); err == nil {
		fmt.Println("10 / 2 =", q)
	}
	if _, err := divide(10, 0); err != nil {
		fmt.Println("divide error:", err)
	}
}
