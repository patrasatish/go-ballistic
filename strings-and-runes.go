package main

import (
	"fmt"
	"unicode/utf8"
)

// rune:
// concept of character in go
// integer that represents a unicode code point

// string:
// equivalent to []byte
// produce the length of raw bytes stored within

// index of string:
// raw bytes at each index

// loop:
// generates hex values of all bytes -> constitute code points

func main() {
	const s = "satish"

	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	// range -> handles string specially
	// decodes each rune along with its offset in string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// explicit handling of rune and offset in string
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// pass rune by value
// 'a', 'b', 'c', 'D' -> rune literals
// can compare rune value with rune literal
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 's' {
		fmt.Println("found s")
	}
}
