package main

import "fmt"

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

func main() {
	var s asciiStr = []byte("test123")
	fmt.Printf("bytes Pre upper: %b\n", s)
	fmt.Printf("string Pre upper: %s\n", s)
	s.ToUpper()
	fmt.Printf("bytes After upper: %b\n", s)
	fmt.Printf("string After upper: %s\n", s)
}
