package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(d int) {
	newNode := &Node{data: d}
	// if head is nil, just add and return (no elements)
	if l.head == nil {
		l.head = newNode
		return
	}
	last := l.head
	// traverse the list to find the end of it
	for last.next != nil {
		last = last.next
	}
	// add the new node at the end of the list
	last.next = newNode
}

func (l *LinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}
	list.append(4)
	list.append(9)
	list.append(21)
	list.display()
}
