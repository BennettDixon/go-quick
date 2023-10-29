package main

import "fmt"

// slices (their array buffers) in go will not be garbage collected
// until all references are gone (evidently); this results
// if we want to improve this and have it be collected, we can
// copy it to a new slice

// in reality maybe this would fetch a file or something
func getBigSlice() []int {
	// this is not a big array but for example
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
}

// does not check edge cases
func getSubSliceNoCopy(start int, end int) []int {
	s := getBigSlice()
	// take some sub slice of the 'big slice'
	newSlice := s[start:end]
	return newSlice
}

// does not check edge cases like start / end vs slice size
func getSubSliceEfficient(start int, end int) []int {
	s := getBigSlice()
	ss := s[start:end]
	// allocate a new slice
	newSlice := make([]int, len(ss))
	// copy data over
	copy(newSlice, ss)
	return newSlice
}

func main() {
	// say we just want the sub slice of our big slice 5:7 (2 elements)
	s := getSubSliceNoCopy(5, 7)
	// now we don't need that big slice, but it is still hanging around
	// since we just have a pointer that references the same buffer
	// let's fix that by using copy
	s = getSubSliceEfficient(5, 7)
	// now we don't have the 'big slice' hanging around anymore
	fmt.Println(s)
}
