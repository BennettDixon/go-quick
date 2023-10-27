package main

import "fmt"

// we could make this a custom type enum, but for simplicity just using ints
func getDayOfWeek(dayOfWeek int) string {
	switch dayOfWeek {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Non Gregorian calendar!"
	}
}

// again purchaseType should probably be some sort of custom type
func printPurchases(pType string, purchases []float32) {
	fmt.Printf("List of %s purchases this week:\n", pType)
	for i := 0; i < len(purchases); i++ {
		d := getDayOfWeek(i + 1)
		fmt.Printf("Your %s on %s was $%.2f\n", pType, d, purchases[i])
	}
}

func main() {
	var weeklyCoffees = []float32{3.25, 4.50, 4.99, 3.25, 4.25, 4.99, 6.50}
	pType := "coffee"
	printPurchases(pType, weeklyCoffees)
	// new line for clarity
	fmt.Println()
	// oops we made a mistake entering the price (scenario), let's update an index!
	weeklyCoffees[0] = 3.49
	weeklyCoffees[1] = 4.99
	printPurchases(pType, weeklyCoffees)
}
