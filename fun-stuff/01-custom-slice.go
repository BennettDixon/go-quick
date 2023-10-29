package main

import "fmt"

// build a custom slice (simple version)

type SliceHeader struct {
	length      int
	zeroElement *int
}

func main() {
	var buffer [512]int
	// same as using a slice [200:205]
	startIndex := 200
	s := SliceHeader{length: 5, zeroElement: &buffer[startIndex]}
	for i := 0; i < s.length; i++ {
		// just insert some 1 values
		v := 1
		buffer[i+startIndex] = v
	}

	// let's add some more elements to our slice
	for i := 0; i < 5; i++ {
		buffer[startIndex+s.length] = 2
		s.length += 1
	}

	// let's check that we actually added the elements as we expect
	// should have 5 1s starting at index 200 then 5 2s after that
	// [... index 200... 1, 1, 1, 1, 1, 2, 2, 2, 2, 2]
	for i := 0; i < s.length; i++ {
		fmt.Printf("element at index %d is %d\n", startIndex+i, buffer[startIndex+i])
	}

	fmt.Println("-----")

	secondStartIndex := 205
	// equivalent of [205:210]
	s2 := SliceHeader{length: 5, zeroElement: &buffer[secondStartIndex]}
	// should just be all 5 2s now
	for i := 0; i < s2.length; i++ {
		fmt.Printf("element at index %d is %d\n", secondStartIndex+i, buffer[secondStartIndex+i])
	}
}
