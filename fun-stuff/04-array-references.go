package main

import "fmt"

// since we pass by value we are actually passing an array in
// which copies it rather than in c where it is a pointer to element 1
func printPassByValueCreatesCopies(arr [4]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

// passing a pointer obviously does not copy the array
// and can be slightly more efficient
func printPassByReferenceNoCopies(arr *[4]int) {
	for i := 0; i < len(*arr); i++ {
		fmt.Printf("%d ", (*arr)[i])
	}
}

func main() {
	a := [4]int{1, 2, 3, 4}
	printPassByValueCreatesCopies(a)
	fmt.Println()
	fmt.Println("-----")
	printPassByReferenceNoCopies(&a)
	fmt.Println()
}
