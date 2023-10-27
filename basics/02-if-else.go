package main

import "fmt"

func main() {
	var name string = "Bennett"
	if len(name) < 6 {
		fmt.Printf("%s is shorter than 6 characters.\n", name)
	} else {
		fmt.Printf("%s is longer than 6 characters.\n", name)
	}
}
