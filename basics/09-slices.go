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
// Assumes purchases are ordered by day of week (dumb assumption but just for example)
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
	coffeePType := "coffee"
	printPurchases(coffeePType, weeklyCoffees)
	// new line for clarity
	fmt.Println()
	// oops we made a mistake entering the price (scenario), let's update an index!
	weeklyCoffees[0] = 3.49
	weeklyCoffees[1] = 4.99
	printPurchases(coffeePType, weeklyCoffees)
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
	// Alright we actually used it Friday & Saturday too!
	// (example appending multiple elements)
	fmt.Println()
	fmt.Println("JK! We used it Friday & Saturday as well!")
	busCommute = append(busCommute, 2.25, 2.25)
	printPurchases(busPType, busCommute)

	// Demonstrate slicing slices & combining slices
	// Let's grab our coffee purchases for Sun-Wed
	fmt.Println()
	sunThruWedCoffee := weeklyCoffees[:4]
	printPurchases(coffeePType, sunThruWedCoffee)

	fmt.Println()
	// This won't print day of the week right but not really the point
	thurThruSatCoffee := weeklyCoffees[4:]
	printPurchases(coffeePType, thurThruSatCoffee)

	// combine slices using dot notation
	fmt.Println()
	allCoffees := append(sunThruWedCoffee, thurThruSatCoffee...)
	printPurchases(coffeePType, allCoffees)
}
