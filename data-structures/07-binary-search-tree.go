package main

import "fmt"

type Node struct {
	Val    int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (t *BinarySearchTree) Insert(val int) {
	if t.Root == nil {
		t.Root = &Node{Val: val}
		return
	}

	next := t.Root
	var prev *Node
	for next != nil {
		// go left if less
		prev = next
		if val < next.Val {
			next = next.Left
		} else {
			// otherwise go right
			next = next.Right
		}
	}
	// once we find a null node, insert
	newNode := &Node{Val: val, Parent: prev}
	if prev.Val < val {
		prev.Right = newNode
	} else {
		prev.Left = newNode
	}
}

func (t *BinarySearchTree) Search(val int) *Node {
	next := t.Root
	for next != nil {
		if next.Val == val {
			return next
		}
		if next.Val < val {
			next = next.Right
		} else {
			next = next.Left
		}
	}
	return nil
}

func (root *Node) DisplayInOrder() {
	if root != nil {
		root.Left.DisplayInOrder()
		fmt.Printf("%d ", root.Val)
		root.Right.DisplayInOrder()
	}
}

func (t *BinarySearchTree) Delete(val int) bool {
	node := t.Search(val)
	if node == nil {
		return false
	}

	isRoot := false
	if node == t.Root {
		isRoot = true
	}

	// Leaf
	if node.Left == nil && node.Right == nil {
		if isRoot {
			t.Root = nil
		} else if node.Parent.Left == node {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
		return true
	}

	// One child; left child is nil
	if node.Left == nil {
		if isRoot {
			t.Root = node.Right
		} else if node.Parent.Left == node {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}
		node.Right.Parent = node.Parent
		return true
	}

	// Once child; right child is nil
	if node.Right == nil {
		if isRoot {
			t.Root = node.Left
		} else if node.Parent.Left == node {
			node.Parent.Left = node.Left
		} else {
			node.Parent.Right = node.Left
		}
		node.Left.Parent = node.Parent
		return true
	}

	// Two children; find the succesor
	succesor := node.Right
	for succesor.Left != nil {
		succesor = succesor.Left
	}
	// swap values
	node.Val, succesor.Val = succesor.Val, node.Val
	var child *Node
	if succesor.Right != nil {
		child = succesor.Right
	}
	if succesor.Parent.Left == succesor {
		succesor.Parent.Left = child
	} else {
		succesor.Parent.Right = child
	}
	if child != nil {
		child.Parent = succesor.Parent
	}
	return true
}

func main() {
	t := BinarySearchTree{}
	t.Insert(1)
	t.Insert(5)
	t.Insert(24452)
	t.Insert(-5)
	t.Insert(2)
	t.Insert(3)
	t.Insert(9)
	t.Root.DisplayInOrder()
	fmt.Println()
	t.Delete(-5)
	t.Delete(5)
	t.Root.DisplayInOrder()
	fmt.Println()
}
