package main

import "fmt"

// fun fact we write left + (right-left)/2 to reduce the risk of overflow
// It is mathematically equivalent to left + right / 2 which is much more intuitive;
// however the risk of overflow is higher
// e.g 2 + 4 = 6 => / 2 = 3
// but 4 - 2 = 2 => / 2 = 1 + 2 = 3 reduces chance of overflow since the highest number is 4 instead of 6
// so if we had 2147000000 elements and were using 32 bit integers
// we could still use 2147000000 - 1073500000 => 1073500000 / 2 = 536750000 + 1073500000 = 1610250000
// Where as we would overflow with the simpler version:
// 2147000000 + 1073500000 => Overflow!

// fun fact the left index is the insertion point if the element is not found
// if we wish to find the insertion point we should return that instead

func binarySearch(sorted []int, target int) int {
	left := 0
	right := len(sorted) - 1
	for left <= right {
		mid := left + (right-left)/2
		if sorted[mid] == target {
			return mid
		}
		// if target is less than mid, move right
		if sorted[mid] < target {
			left = mid + 1
		} else {
			// if it is greater than target, set window to left side
			right = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 3, 50, 64, 75, 102}
	i := binarySearch(list, 999)
	if i != -1 {
		fmt.Println("Test case failed! Should not find index for 999 it is not present")
	}
	i = binarySearch(list, 3)
	if i != 2 {
		fmt.Println("Test case failed! i should equal 2 for search of 3.")
	}
	i = binarySearch(list, 64)
	if i != 4 {
		fmt.Println("Test case failed! i should equal 2 for search of 64")
	}
	// test empty list
	list = []int{}
	i = binarySearch(list, 2)
	if i != -1 {
		fmt.Println("Should not find an index for an empty list.")
	}
	fmt.Println("If nothing else is shown above, all indexes were found as expected.")
}
