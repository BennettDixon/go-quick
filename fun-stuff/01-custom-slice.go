package main

import "fmt"

// build a custom slice (simple version)

type SliceHeader struct {
	Length      int
	ZeroElement *int
	ZeroIndex   int
	Capacity    int
}

func (s SliceHeader) remainingCapacity() int {
	return s.Capacity - s.Length
}

func main() {
	var buffer [512]int
	s := SliceHeader{Length: 0, ZeroElement: &buffer[0], ZeroIndex: 0, Capacity: 512}
	for i := 0; i < 5; i++ {
		// just insert some 1 values
		buffer[s.Length+s.ZeroIndex] = 1
		s.Length += 1
	}

	// let's add some more elements to our slice
	for i := 0; i < 5; i++ {
		// insert some 2s after the 1s
		buffer[s.ZeroIndex+s.Length] = 2
		s.Length += 1
	}

	// let's check that we actually added the elements as we expect
	// should have 5 1s starting at index 0 then 5 2s after that
	// [1, 1, 1, 1, 1, 2, 2, 2, 2, 2]
	for i := s.ZeroIndex; i < s.Length+s.ZeroIndex; i++ {
		fmt.Printf("element at index of first slice %d is %d\n", i, buffer[i])
	}

	fmt.Println("-----")

	// equivalent of [5:15]
	startIndex := 5
	s2 := SliceHeader{Length: 10, ZeroElement: &buffer[startIndex], ZeroIndex: startIndex, Capacity: s.Capacity - startIndex}
	// should just be 2s for the first 5 & 0 after that
	// [2, 2, 2, 2, 2, 0, 0, 0, 0, 0]
	for i := s2.ZeroIndex; i < s2.Length+s2.ZeroIndex; i++ {
		fmt.Printf("element at index of new slice %d is %d\n", i, buffer[i])
	}
}
