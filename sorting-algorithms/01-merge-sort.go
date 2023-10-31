package main

import "fmt"

func mergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	m := len(s) / 2
	l := s[:m]
	r := s[m:]
	// recursively stack all of these calls until base case is hit
	l = mergeSort(l)
	r = mergeSort(r)
	// finally sort the elements
	i, j, k := 0, 0, 0
	res := make([]int, len(s))
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			res[k] = l[i]
			i++
		} else {
			res[k] = r[j]
			j++
		}
		k++
	}
	// add any remaining elements
	for i < len(l) {
		res[k] = l[i]
		k++
		i++
	}
	for j < len(r) {
		res[k] = r[j]
		k++
		j++
	}
	return res
}

func main() {
	s1 := []int{1, 10, 5, 20, 3, 924, 12, 12, 1, 342, 421}
	fmt.Println(s1)
	fmt.Println("Sorting with merge sort...")
	s1 = mergeSort(s1)
	fmt.Println(s1)
	s2 := []int{}
	s2 = mergeSort(s2)
	fmt.Println("Empty list sorted.")
}
