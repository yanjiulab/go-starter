package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Strings are immutable UTF-8 byte sequences.
	text := "Hello, 世界"
	fmt.Println("bytes:", []byte(text))
	fmt.Println("runes:", []rune(text))
	fmt.Println("first byte:", text[0], "first rune:", []rune(text)[0])
	fmt.Println("contains 世界?", strings.Contains(text, "世界"))
	fmt.Println("byte length:", len(text), "rune count:", utf8.RuneCountInString(text))

	// range iterates over runes (Unicode code points).
	fmt.Println("iterate by rune:")
	for i, r := range text {
		fmt.Printf("  index=%d rune=%c\n", i, r)
	}

	replaced := strings.ReplaceAll(text, "Hello", "Hi")
	fmt.Println("replace:", replaced, "upper:", strings.ToUpper(replaced))
	fmt.Println("split:", strings.Split(replaced, ", "))
}
