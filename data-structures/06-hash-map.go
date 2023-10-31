package main

import "fmt"

type Node struct {
	Key   string
	Value int
	Next  *Node
}

type HashTable struct {
	Table []*Node
	Size  int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Table: make([]*Node, size),
		Size:  size,
	}
}

func (h *HashTable) hashFunc(key string) int {
	sum := 0
	// generate the index based on the key
	for _, v := range key {
		// get the value of the byte and add to sum
		sum += int(v)
	}
	// return an index based on remainder of key sum % size
	return sum % h.Size
}

func (h *HashTable) Put(key string, value int) {
	index := h.hashFunc(key)
	newNode := &Node{Key: key, Value: value}
	if h.Table[index] == nil {
		// just add the node if it doesn't exist
		h.Table[index] = newNode
	} else {
		// if it exists we need to handle the collision
		curNode := h.Table[index]
		// loop to end of list and add it on the end
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = newNode
	}
}

func (h *HashTable) Get(key string) (int, bool) {
	index := h.hashFunc(key)
	curNode := h.Table[index]
	// loop through the list and find the matching key
	for curNode != nil {
		if curNode.Key == key {
			return curNode.Value, true
		}
		curNode = curNode.Next
	}
	return 0, false
}

func main() {
	hashTable := NewHashTable(10)
	hashTable.Put("key1", 10)
	hashTable.Put("key2", 20)

	if value, found := hashTable.Get("key1"); found {
		fmt.Println("Found key1:", value)
	}

	if _, found := hashTable.Get("key3"); !found {
		fmt.Println("key3 not found")
	}
}
