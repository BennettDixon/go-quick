package main

import "fmt"

func main() {
	fmt.Println("Printing I 0-10!")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 10
	fmt.Println("\nPrinting X 10-0!")
	for x >= 0 {
		fmt.Println(x)
		x -= 1
	}

	fmt.Println("\nPrinting Y 0-5!")
	var y int = 0
	for {
		if y > 5 {
			fmt.Println("Break!")
			break
		}
		fmt.Println(y)
		y++
	}
}
