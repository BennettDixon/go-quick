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

func main() {
	var weeklyCoffees = [7]float32{3.25, 4.50, 4.99, 3.25, 4.25, 4.99, 6.50}
	fmt.Println("List of coffee purchases this week:")
	for i := 0; i < len(weeklyCoffees); i++ {
		d := getDayOfWeek(i + 1)
		fmt.Printf("Your coffee on %s was $%.2f\n", d, weeklyCoffees[i])
	}
}
