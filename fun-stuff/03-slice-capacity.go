package main

import (
	"errors"
	"fmt"
)

func willPanic(slice []int, value int) (newSlice []int, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("Encountered error")
		}
	}()
	// try to grow our slice by 1 in a poor way
	// will panic because capacity is len
	newSlice = slice[:len(slice)+1]
	return slice, nil
}

func addCapacityByValueNoReturnWillNotWork(slice []int, value int) {
	// use make to grow our capacity by value
	slice = make([]int, len(slice), len(slice)+value)
}

// prefered way
func addCapacityByValue(slice []int, value int) []int {
	// use make to grow our capacity by value
	return make([]int, len(slice), len(slice)+value)
}

// needlessly complex but just an interesting test
func addCapacityByReference(slice *[]int, value int) {
	*slice = make([]int, len(*slice), len(*slice)+value)
}

func main() {
	s := []int{1, 2, 3}
	_, err := willPanic(s, 1)
	if err != nil {
		fmt.Println("Paniced adding new capacity to slice in poor way as expected.")
	}

	fmt.Printf("by value broken: Capacity pre insert %d\n", cap(s))
	addCapacityByValueNoReturnWillNotWork(s, 5)
	fmt.Printf("by value broken: Capacity after insert (should be same) %d\n", cap(s))
	fmt.Println("-----------")
	fmt.Printf("by value: Capacity pre insert %d\n", cap(s))
	newSlice := addCapacityByValue(s, 5)
	fmt.Printf("by value: Capacity after insert (should be %d) %d\n", len(s)+5, cap(newSlice))
	fmt.Println("-----------")
	fmt.Printf("by value: Capacity pre insert %d\n", cap(s))
	addCapacityByReference(&newSlice, 10)
	fmt.Printf("by value: Capacity after insert (should be %d) %d\n", len(newSlice)+10, cap(newSlice))
}
