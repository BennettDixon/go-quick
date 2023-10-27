package main

import "fmt"

// pointers are cool, brings me back to C :D
func main() {
	a := 1
	b := a
	b += 1
	fmt.Printf("Copies (no pointers), a: %d, b: %d\n", a, b)
	if &a == &b {
		fmt.Println("The copies shouldn't have the same address, what is going on!?!?")
	} else {
		fmt.Println("As expected, the copies have different addresses!")
		fmt.Printf("Memory address of a is %d, b is %d\n", &a, &b)
	}

	// New line for clarity of copies vs pointers
	fmt.Println()

	c := 1
	var d *int = &c
	c += 1
	// increment them both via the pointer
	*d += 1
	fmt.Printf("Pointers (equal after-add since same reference)! c: %d, d: %d\n", c, *d)
	fmt.Printf("Memory address (same reference / address)! c: %d, d: %d\n", &c, d)
	if &c != d {
		fmt.Println("Memory address should be the same, what is going on this isn't c! We can't increment it!!")
	} else {
		fmt.Println("Memory addresses are the same; thank google for go!")
	}
}
