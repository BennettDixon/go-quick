package main

import (
	"errors"
	"fmt"
)

// Setup our Error struct & ErrorType
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

// Simple add one by reference function
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

func willPanic() (err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("My panic error")
		}
	}()
	// try to grow our slice by 1 in a poor way
	// will panic because capacity is len
	panic("panic!")
}

// Go doesn't have exceptions or try / catch
// Use multiple return values instead, or single in this case
// Convention is left return value, right error value if using multiples (res, err)
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

	// Show handling a panic
	basicErr := willPanic()
	if basicErr != nil {
		fmt.Println("Handled panic and got error back")
		fmt.Println(basicErr.Error())
	}
}
