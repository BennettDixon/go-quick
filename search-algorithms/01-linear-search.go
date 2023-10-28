package main

import "fmt"

func linearSearch(l []int, target int) int {
	// just loop and find it
	for i := 0; i < len(l); i++ {
		if l[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 3, 50, 64, 75, 102}
	i := linearSearch(list, 999)
	if i != -1 {
		fmt.Println("Test case failed! Should not find index for 999 it is not present")
	}
	i = linearSearch(list, 3)
	if i != 2 {
		fmt.Println("Test case failed! i should equal 2 for search of 3.")
	}
	i = linearSearch(list, 64)
	if i != 4 {
		fmt.Println("Test case failed! i should equal 2 for search of 64")
	}
	// test empty list
	list = []int{}
	i = linearSearch(list, 2)
	if i != -1 {
		fmt.Println("Should not find an index for an empty list.")
	}
	fmt.Println("If nothing else is shown above, all indexes were found as expected.")
}
