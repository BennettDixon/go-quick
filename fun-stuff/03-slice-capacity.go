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
	newSlice = slice[:len(slice)+value]
	return
}

func addCapacityByValueNoReturnWillNotWork(slice []int, value int) {
	// use make to grow our capacity by value
	slice = make([]int, len(slice), len(slice)+value)
}

// simple custom copy that doesn't take into all edge cases
func customCopy(dst []int, src []int) {
	// handle if capacity of destination can't handle src size
	if cap(dst) < len(src) {
		dst = make([]int, len(src))
	}
	for i, e := range src {
		dst[i] = e
	}
}

func addCapacityNoCopy(slice []int, value int) []int {
	return make([]int, len(slice), cap(slice)+value)
}

// prefered way
func addCapacityByValue(slice []int, value int) []int {
	newSlice := addCapacityNoCopy(slice, value)
	// copy our values with a loop
	customCopy(newSlice, slice)
	return newSlice
}

// needlessly complex but just an interesting test
func addCapacityByReference(slice *[]int, value int) {
	newSlice := addCapacityNoCopy(*slice, value)
	customCopy(newSlice, *slice)
	*slice = newSlice
}

// built in copy method is preferred
func addCapactyByValueUsingCopy(slice []int, value int) []int {
	newSlice := addCapacityNoCopy(slice, value)
	copy(newSlice, slice)
	return newSlice
}

func printSlice(slice []int) {
	for _, e := range slice {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}

func main() {
	s := []int{1, 2, 3}
	_, err := willPanic(s, 1)
	if err != nil {
		fmt.Println("Paniced adding new capacity to slice in poor way as expected.")
	}

	fmt.Println("----by-value-broken-------")
	fmt.Printf("Length pre insert %d\n", len(s))
	fmt.Printf("Capacity pre insert %d\n", cap(s))
	addCapacityByValueNoReturnWillNotWork(s, 5)
	fmt.Printf("Length after insert %d\n", len(s))
	fmt.Printf("Capacity after insert (should be same) %d\n", cap(s))
	printSlice(s)
	fmt.Println("----by-value-------")
	fmt.Printf("Length pre insert %d\n", len(s))
	fmt.Printf("Capacity pre insert %d\n", cap(s))
	newSlice := addCapacityByValue(s, 5)
	fmt.Printf("Length after insert %d\n", len(newSlice))
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(s)+5, cap(newSlice))
	printSlice(newSlice)
	fmt.Println("-----by reference------")
	fmt.Printf("Length pre insert %d\n", len(newSlice))
	fmt.Printf("Capacity pre insert %d\n", cap(newSlice))
	oldCapacity := cap(newSlice)
	addCapacityByReference(&newSlice, 10)
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(newSlice), oldCapacity+10)
	fmt.Printf("Length after insert %d\n", len(newSlice))
	printSlice(newSlice)
	fmt.Println("----by-value-built-in-copy-------")
	fmt.Printf("Length pre insert %d\n", len(newSlice))
	fmt.Printf("Capacity pre insert %d\n", cap(newSlice))
	copiedSlice := addCapactyByValueUsingCopy(newSlice, 10)
	printSlice(copiedSlice)
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(newSlice)+10, cap(copiedSlice))
	fmt.Printf("Length after insert %d\n", len(copiedSlice))
}
