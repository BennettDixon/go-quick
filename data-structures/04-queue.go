package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) dequeue() *int {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return &item
}

func (q *Queue) display() {
	for i := 0; i < len(q.items); i++ {
		fmt.Printf("%d ", q.items[i])
	}
	fmt.Println()
}

func main() {
	q := Queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.display()
	firstItem := q.dequeue()
	for firstItem != nil {
		fmt.Printf("Item: %d\n", *firstItem)
		firstItem = q.dequeue()
	}
}
