package main

import "fmt"

func printName(name string) {
	fmt.Printf("Your name is %s", name)
}

func addNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func combineNames(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func main() {
	// Type inference here instead :D
	firstName := "Bennett"
	lastName := "Dixon"
	// add the names together
	name := combineNames(firstName, lastName)
	// print the name
	printName(name)
	// print a newline
	fmt.Println()
	num1 := 1
	num2 := 2
	target := 3
	result := addNumbers(num1, num2)
	if target != result {
		fmt.Println("Math doesn't work, this is not good!")
	} else {
		fmt.Println("Math works, things are fine in life!")
	}
}
