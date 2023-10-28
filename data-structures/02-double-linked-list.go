package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data int
}

type DoublyLinkedList struct {
	head *Node
}

func (l *DoublyLinkedList) append(d int) {
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
	newNode.prev = last
}

func (l *DoublyLinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (l *DoublyLinkedList) displayReverse() {
	last := l.head
	// Find the end of the list
	for last.next != nil {
		last = last.next
	}
	// Now print it in reverse order
	for last != nil {
		fmt.Printf("%d ", last.data)
		last = last.prev
	}
	fmt.Println()
}

func main() {
	list := DoublyLinkedList{}
	list.append(4)
	list.append(9)
	list.append(21)
	list.display()
	list.displayReverse()
}
