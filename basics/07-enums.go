package main

type Suit int

const (
	Spade = iota + 1
	Diamond
	Club
	Heart
)

type Card struct {
	suit Suit
	num  int
}

func (c *Card) PrintSuit() {
	switch c.suit {
	case Spade:
		print("It is a spade!")
	case Diamond:
		print("It is a diamond!")
	case Club:
		print("It is a club!")
	case Heart:
		print("It is a heart!")
	default:
		print("Unknown Suit! This is cards not chess!")
	}
}

func main() {
	c := Card{suit: Diamond, num: 12}
	c.PrintSuit()
	print("\n")
}
