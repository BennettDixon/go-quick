package main

import "fmt"

func main() {
	// Type inference here instead :D
	firstName := "Bennett"
	lastName := "Dixon"
	useLastName := true
	var name string
	if useLastName {
		name = lastName
	} else {
		name = firstName
	}
	switch name {
	case "Bennett":
		fmt.Println("You used your first name.")
	case "Dixon":
		fmt.Println("You used your last name.")
	}
}
