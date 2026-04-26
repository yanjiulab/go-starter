package main

import "fmt"

// counter captures outer variable and keeps state.
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// makeAdder is a closure factory.
func makeAdder(base int) func(int) int {
	return func(v int) int {
		return base + v
	}
}

func main() {
	// State is isolated per closure instance.
	c := counter()
	fmt.Println(c(), c(), c())

	add10 := makeAdder(10)
	add100 := makeAdder(100)
	fmt.Println("add10(5):", add10(5))
	fmt.Println("add100(5):", add100(5))
}
