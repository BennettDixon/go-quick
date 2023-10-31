package main

import "fmt"

// Quick sort is better for space than merge sort; O(log(n)) instead of O(n)
// Otherwise merge sort is slightly better time complexity since worst case is still O(n log(n))
// O(log(n)) Space complexity
// O(n log(n)) Time complexity BEST / AVERAGE
// O(n^2) Time complexity WORST

func quickSort(s []int, l int, h int) {
	// base case
	if l >= h {
		return
	}
	p := s[h]
	i, j := l, h
	for i < j {
		// find left swap spot
		for s[i] <= p && i < j {
			i++
		}
		// find right swap spot
		for s[j] >= p && j > i {
			j--
		}
		// now we need to swap then continue until we hit the same spot
		swap(s, i, j)
	}
	// finally swap with the pivot
	swap(s, i, h)

	quickSort(s, l, i-1)
	quickSort(s, i+1, h)
}

func swap(s []int, i int, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func main() {
	s1 := []int{1, 45, 3, 2, 49124, 4, 12424, 99, 424, 98, 1, 2, 2, 33, 4242}
	fmt.Println(s1)
	quickSort(s1, 0, len(s1)-1)
	fmt.Println("Sorted!")
	fmt.Println(s1)
}
