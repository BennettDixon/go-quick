package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	myPerson := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       81,
	}
	fmt.Printf("%s, %s is %d years old.\n", myPerson.firstName, myPerson.lastName, myPerson.age)
	fmt.Println("They just had a birthday!")
	myPerson.age += 1
	fmt.Printf("Now they are %d years old.\n", myPerson.age)
}
