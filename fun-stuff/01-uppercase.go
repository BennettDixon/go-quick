package main

import (
	"fmt"
	"unicode"
)

// rune = int32 in go
type uniStr []rune

// byte is just byte
type asciiStr []byte

// modifying underlying array so don't need to pass a pointer to slice
func (str asciiStr) ToUpper() {
	var offset byte = 'a' - 'A'
	// should be 32 (constant offset for ascii)
	for i, e := range str {
		// ascii letters a-z are higher than A-Z (note)
		// make sure it is a lowercase letter
		if e >= 'a' && e <= 'z' {
			// add the offset of capitals which is always constant
			str[i] = e - offset
		}
	}
}

func (str uniStr) ToUpper() {
	for i, e := range str {
		str[i] = unicode.ToUpper(e)
	}
}

func main() {
	var s asciiStr = []byte("test123")
	fmt.Println("-------Asci------")
	fmt.Printf("bytes pre upper: %b\n", s)
	fmt.Printf("string pre upper: %s\n", s)
	s.ToUpper()
	fmt.Printf("bytes after upper: %b\n", s)
	fmt.Printf("string after upper: %s\n", s)
	fmt.Println("-----Unicode-----")
	var helloS uniStr = []rune("Hello, 世界")
	var emojiS uniStr = []rune("I love 🌞")
	fmt.Printf("runes pre upper: %#U, %#U\n", helloS, emojiS)
	fmt.Printf("string pre upper: helloS: %s\nemojiS: %s\n", string(helloS), string(emojiS))
	helloS.ToUpper()
	emojiS.ToUpper()
	fmt.Printf("runes after upper: %#U, %#U\n", helloS, emojiS)
	fmt.Printf("string after upper: helloS: %s\nemojiS: %s\n", string(helloS), string(emojiS))
}
