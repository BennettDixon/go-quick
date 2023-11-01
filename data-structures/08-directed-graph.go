package main

import "fmt"

type DirectedGraph map[string][]string

// add an edge to our adjacency list
func addEdge(graph DirectedGraph, u, v string) {
	graph[u] = append(graph[u], v)
}

// for depth first traversal we use a stack
// DEPTH FIRST USES STACKS BUT STAy going down! (jokes)
// Recursively print elements depth first
func (d DirectedGraph) PrintDepthFirstRecursion(s string, v map[string]bool) {
	v[s] = true
	fmt.Println(s)
	for _, n := range d[s] {
		if !v[n] {
			d.PrintDepthFirstRecursion(n, v)
		}
	}
}

// Iteratively print elements depth first
func (d DirectedGraph) PrintDepthFirstIterative(s string) {
	// track nodes visited already
	v := make(map[string]bool)
	// make a stack
	stack := []string{}
	// push the root onto the stack
	stack = append(stack, s)
	for len(stack) > 0 {
		// pop an item off the stack
		i := len(stack) - 1
		element := stack[i]
		stack = stack[:i]
		// loop through the neighbors
		if !v[element] {
			fmt.Println(element)
			v[element] = true
			// Add unvisited neighbors to stack
			for _, neighbor := range d[element] {
				if !v[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
}

// for breadth first traversal we use a queue
// BREADTH FIRST USES QUEUES BUT QUE bread?? (jokes)
func (d DirectedGraph) PrintBreadthFirst(s string) {
	// the logic is the same as DFS, but we just use a queue
	v := make(map[string]bool)
	queue := []string{}
	queue = append(queue, s)
	for len(queue) > 0 {
		// dequeue the element
		element := queue[0]
		queue = queue[1:]
		// if we haven't visited the element, print & add unvisited neighbors
		if !v[element] {
			fmt.Println(element)
			v[element] = true
			for _, n := range d[element] {
				// only add the neighbor if we haven't visited it
				if !v[n] {
					queue = append(queue, n)
				}
			}
		}
	}
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

	// they actually won't match because recursion reverses
	// the output, but if we want to match we could
	// instead add neighbors in reverse in our iterative approach
	fmt.Println("Depth first recursion:")
	// make a map to know if we visited the nodes since we use recursion
	dFirstVisited := make(map[string]bool)
	dg.PrintDepthFirstRecursion("Collect", dFirstVisited)
	// iterative will make this map itself
	fmt.Println("Depth first iterative")
	dg.PrintDepthFirstIterative("Collect")
	fmt.Println("Breadth first:")
	dg.PrintBreadthFirst("Collect")
}
