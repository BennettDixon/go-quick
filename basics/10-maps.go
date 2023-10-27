package main

import "fmt"

// again items should be a custom type or struct, just simple maps for this one though
// also should use functions etc. to reduce duplication, but again just maps for this one
func main() {
	itemPricing := make(map[string]float32)
	lampKey := "Lamp"
	jacketKey := "Jacket"
	itemPricing[jacketKey] = 44.99
	itemPricing[lampKey] = 29.99
	fmt.Printf("A lamp costs $%.2f\n", itemPricing[lampKey])
	fmt.Printf("A jacket costs $%.2f\n", itemPricing[jacketKey])
	shoppingCart := make(map[string]int)
	shoppingCart[jacketKey] = 0
	shoppingCart[lampKey] = 0
	// let's buy one jacket
	shoppingCart[jacketKey] += 1
	// and two lamps, it is dark!!
	shoppingCart[lampKey] += 2
	lampQuantity, lampExists := shoppingCart[lampKey]
	// can check if a key exists in the map using a secondary accessor
	if lampExists {
		fmt.Println("The lamp key exists!")
	}
	fmt.Printf(
		"You have %d lamp(s) that cost %.2f, and %d jacket(s) that cost %.2f in your cart.\n",
		lampQuantity,
		itemPricing[lampKey],
		shoppingCart[jacketKey],
		itemPricing[jacketKey],
	)
}
