package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) pop() *int {
	l := len(s.items)
	if l == 0 {
		return nil
	}
	lastIndex := l - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return &item
}

func (s *Stack) display() {
	for i := 0; i < len(s.items); i++ {
		fmt.Printf("%d ", s.items[i])
	}
	fmt.Println()
}

func main() {
	s := Stack{}
	nilPtr := s.pop()
	if nilPtr == nil {
		fmt.Println("Pop returns nil as expected when no elements in stack")
	}
	s.push(1)
	s.push(4)
	s.push(5)
	s.display()
	var item *int = s.pop()
	for item != nil {
		fmt.Printf("Popped item %d\n", *item)
		item = s.pop()
	}
}
