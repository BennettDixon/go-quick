package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const nihongo = "日本語"
	// one 'gotcha' in go is that the index of the bytes
	// for unicode characters is not necessarily the iteration of the loop
	// for example these indexes will be 0, 3, 6
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	// if we need something like this, use the std unicode library
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
