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
// Converted to use slices now for flexibility!
// turns out not all purchases are weekly based!
func printPurchases(pType string, purchases []float32) {
	fmt.Printf("List of %s purchases this week:\n", pType)
	for i := 0; i < len(purchases); i++ {
		d := getDayOfWeek(i + 1)
		fmt.Printf("Your %s on %s was $%.2f\n", pType, d, purchases[i])
	}
}

func main() {
	// Set length of array
	var weeklyCoffees = []float32{3.25, 4.50, 4.99, 3.25, 4.25, 4.99, 6.50}
	pType := "coffee"
	printPurchases(pType, weeklyCoffees)
	// new line for clarity
	fmt.Println()
	// oops we made a mistake entering the price (scenario), let's update an index!
	weeklyCoffees[0] = 3.49
	weeklyCoffees[1] = 4.99
	printPurchases(pType, weeklyCoffees)
	// new line for clarity
	fmt.Println()
	// Let go figure out length of array
	// Hybrid work, woohoo! Only used the bus 3 times this week! (scenario)
	busPType := "bus"
	var busCommute = []float32{2.25, 2.25, 2.25}
	printPurchases(busPType, busCommute)
	// new line for clarity
	fmt.Println()
	fmt.Println("Oops turned out we needed to go in on Thursday too!")
	// Add a new purchase to the slice
	busCommute = append(busCommute, 2.25)
	printPurchases(busPType, busCommute)
}
