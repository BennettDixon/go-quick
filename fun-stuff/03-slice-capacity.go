package main

import (
	"errors"
	"fmt"
)

// None of this is very practical but just interesting
// to learn how slices work a bit better
// roughly following https://go.dev/blog/slices

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

func Insert(slice []int, index, value int) (newSlice []int, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("Encountered error inserting element into slice!")
		}
	}()
	// Grow the slice by one element.
	newSlice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	newSlice[index] = value
	// Return the result.
	return
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Append appends the items to the slice.
// First version: just loop calling Extend.
func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

// Append appends the elements to the slice.
// Efficient version: no more than one allocation for new slice
func AppendEfficient(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

func main() {
	s := []int{1, 2, 3}
	_, err := willPanic(s, 1)
	if err != nil {
		fmt.Println("Paniced adding new capacity to slice in poor way as expected.")
	}
	// my stuff; messing with slices
	fmt.Println("----by-value-broken-------")
	fmt.Printf("Length pre insert %d\n", len(s))
	fmt.Printf("Capacity pre insert %d\n", cap(s))
	addCapacityByValueNoReturnWillNotWork(s, 5)
	fmt.Printf("Length after insert %d\n", len(s))
	fmt.Printf("Capacity after insert (should be same) %d\n", cap(s))
	fmt.Println(s)

	fmt.Println("----by-value-------")
	fmt.Printf("Length pre insert %d\n", len(s))
	fmt.Printf("Capacity pre insert %d\n", cap(s))
	newSlice := addCapacityByValue(s, 5)
	fmt.Printf("Length after insert %d\n", len(newSlice))
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(s)+5, cap(newSlice))
	fmt.Println(newSlice)

	fmt.Println("-----by reference------")
	fmt.Printf("Length pre insert %d\n", len(newSlice))
	fmt.Printf("Capacity pre insert %d\n", cap(newSlice))
	oldCapacity := cap(newSlice)
	addCapacityByReference(&newSlice, 10)
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(newSlice), oldCapacity+10)
	fmt.Printf("Length after insert %d\n", len(newSlice))
	fmt.Println(newSlice)

	fmt.Println("----by-value-built-in-copy-------")
	fmt.Printf("Length pre insert %d\n", len(newSlice))
	fmt.Printf("Capacity pre insert %d\n", cap(newSlice))
	copiedSlice := addCapactyByValueUsingCopy(newSlice, 10)
	fmt.Println(copiedSlice)
	fmt.Printf("Capacity after insert (should be %d) %d\n", cap(newSlice)+10, cap(copiedSlice))
	fmt.Printf("Length after insert %d\n", len(copiedSlice))

	fmt.Println("------test-insert-------")
	// following examples in the blog
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)
	// add some capacity to allow for our insert
	// could have done in the make of slice, but this is a better demo
	slice = addCapacityByValue(slice, 10)
	slice, err = Insert(slice, 5, 99)
	if err != nil {
		fmt.Println("Error occured inserting element to slice! Shouldn't happen since we set capacity > len")
		fmt.Println(err)
	}
	fmt.Println(slice)

	fmt.Println("-----Extend example-------")
	slice = make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice = Extend(slice, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
		fmt.Println("address of 0th element:", &slice[0])
	}

	fmt.Println("-----append example-------")
	slice = []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	slice = Append(slice, 5, 6, 7, 8)
	fmt.Println(slice)

	fmt.Println("-----append-explode-example-------")
	slice = []int{0, 1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8, 9}
	fmt.Println(slice)
	fmt.Println(slice2)
	slice = Append(slice, slice2...)
	fmt.Println(slice)

	fmt.Println("-----append-efficient-example-----")
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 = []int{55, 66, 77}
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1 = AppendEfficient(slice1, slice2...)
	fmt.Println(slice1)

	// built in append is generic
	fmt.Println("-----built-in-append-generic-example-----")
	// Create a couple of starter slices.
	slice = []int{1, 2, 3}
	slice2 = []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	// Make a copy of a slice (of int).
	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	// Copy a slice to the end of itself.
	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)
}
