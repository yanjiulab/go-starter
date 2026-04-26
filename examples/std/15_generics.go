package main

import "fmt"

// Number is a type constraint.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum works for all Number-constrained types.
func Sum[T Number](vals []T) T {
	var total T
	for _, v := range vals {
		total += v
	}
	return total
}

// Map is a generic transform helper.
func Map[T any, R any](vals []T, fn func(T) R) []R {
	out := make([]R, 0, len(vals))
	for _, v := range vals {
		out = append(out, fn(v))
	}
	return out
}

func main() {
	fmt.Println("sum int:", Sum([]int{1, 2, 3}))
	fmt.Println("sum int64:", Sum([]int64{10, 20, 30}))
	fmt.Println("sum float:", Sum([]float64{1.5, 2.5}))

	squared := Map([]int{1, 2, 3, 4}, func(v int) int { return v * v })
	fmt.Println("mapped squared:", squared)
}
