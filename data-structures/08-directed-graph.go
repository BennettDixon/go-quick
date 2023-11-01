package main

import "fmt"

type DirectedGraph map[string][]string

// FIFO (Breadth first)
type GraphQueue []*DirectedGraph

// LIFO (Depth first)
type GraphStack []*DirectedGraph

// add element to end of queue
func (q GraphQueue) Enqueue(d *DirectedGraph) {
	q = append(q, d)
}

// remove first element from start and reallocate
func (q GraphQueue) Dequeue() *DirectedGraph {
	if len(q) == 0 {
		return nil
	}
	d := q[0]
	q = q[1:]
	return d
}

// push to end of stack
func (s GraphStack) Push(d *DirectedGraph) {
	s = append(s, d)
}

// remove from end of stack
func (s GraphStack) Pop() *DirectedGraph {
	i := len(s) - 1
	d := s[i]
	s = s[:i]
	return d
}

// add an edge to our adjacency list
func addEdge(graph DirectedGraph, u, v string) {
	graph[u] = append(graph[u], v)
}

// for depth first traversal we use a stack
// DEPTH FIRST USES STACKS BUT STAy going down! (jokes)
func (d *DirectedGraph) PrintDepthFirst() {

}

// for breadth first traversal we use a queue
// BREADTH FIRST USES QUEUES BUT QUE bread?? (jokes)
func (d *DirectedGraph) PrintBreadthFirst() {

}

func main() {
	dg := make(DirectedGraph)
	// 	  Collect
	// 	  /    \
	// Backup  Preprocess
	// 	      /       \
	//   Validate--->Train
	// 		\       /
	//    	Evaluate
	addEdge(dg, "Collect", "Preprocess")
	addEdge(dg, "Collect", "Backup")
	addEdge(dg, "Preprocess", "Train")
	addEdge(dg, "Preprocess", "Validate")
	addEdge(dg, "Train", "Evaluate")
	addEdge(dg, "Validate", "Evaluate")

	fmt.Println("Breadth first:")
	dg.PrintBreadthFirst()
	fmt.Println("Depth first:")
	dg.PrintDepthFirst()
}
