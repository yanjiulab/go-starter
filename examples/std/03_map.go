package main

import "fmt"

func main() {
	// nil map supports read but panics on write.
	var nilMap map[string]int
	fmt.Println("read nil map:", nilMap["x"])

	// Literal initialization + update.
	m := map[string]int{"go": 1}
	m["lang"] = 2
	m["go"]++

	v, ok := m["go"]
	fmt.Println("exists?", ok, "value:", v)
	fmt.Println("missing key default value:", m["missing"])

	delete(m, "lang")

	// Range order is intentionally random in Go maps.
	for k, val := range m {
		fmt.Println(k, val)
	}

	// clear removes all entries but keeps map usable.
	clear(m)
	fmt.Println("after clear len:", len(m))
}
