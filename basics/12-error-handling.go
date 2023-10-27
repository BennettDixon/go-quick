package main

import "fmt"

type ErrorType int

const (
	OutOfBoundsError = iota + 1
	NilPointerError
)

type Error struct {
	msg string
	t   ErrorType
}

func (e *Error) Print() {
	fmt.Printf("Encountered error of type %d, message: %s\n", e.t, e.msg)
}

// pretty bad function but just for demo of how you might handle some errors
func addOneToPointer(value *int) *Error {
	obError := Error{msg: "Out of bounds value encountered! Cannot add past 10.", t: OutOfBoundsError}
	nilPtrError := Error{msg: "Nil pointer error encountered!", t: NilPointerError}
	if value == nil {
		return &nilPtrError
	}
	if *value >= 10 {
		return &obError
	}
	*value += 1
	return nil
}

// Go doesn't have exceptions or try / catch
// Use multiple return values instead
// Convention is left return value, right error value
func main() {
	// Demo nil ptr error
	err := addOneToPointer(nil)
	if err != nil {
		err.Print()
	}

	// Demo OB error
	a := 1
	for {
		err := addOneToPointer(&a)
		if err != nil {
			// just print out the error and break for now
			err.Print()
			break
		}
	}
	fmt.Printf("Added a to its max capacity, result: %d\n", a)
}
