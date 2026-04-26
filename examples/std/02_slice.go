package main

import "fmt"

func main() {
	// nil slice vs empty slice.
	var nilSlice []int
	empty := []int{}
	fmt.Println("nil len/cap:", len(nilSlice), cap(nilSlice))
	fmt.Println("empty len/cap:", len(empty), cap(empty))

	// make creates a slice with specific capacity.
	s := make([]int, 0, 2)
	s = append(s, 1, 2, 3)
	s = append(s, 4, 5) // may trigger reallocation

	// Slicing shares the same underlying array.
	base := []int{10, 20, 30, 40}
	sub := base[1:3]
	sub[0] = 99

	dst := make([]int, len(sub))
	copy(dst, sub)
	dst[0] = 777

	// Deep copy shortcut using append.
	s2 := append([]int(nil), s...)

	fmt.Println("nil?", nilSlice == nil, "empty nil?", empty == nil)
	fmt.Println("s len/cap:", len(s), cap(s), "s2:", s2)
	fmt.Println("base:", base, "sub(shared):", sub, "copy(independent):", dst)
	for i, v := range s {
		fmt.Println("range s:", i, v)
	}
}
