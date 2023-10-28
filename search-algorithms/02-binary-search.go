package main

import "fmt"

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
	i := binarySearch(list, 3)
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
	fmt.Println("All indexes found as expected.")
}
