package main

import "fmt"

type UndirectedGraph map[string][]string

// add an edge to our adjacency list
func addEdge(graph UndirectedGraph, u, v string) {
	// undirected graphs are bi-directional; add bi-direction edges
	graph[u] = append(graph[u], v)
	graph[v] = append(graph[v], u)
}

// The only difference in un-direct vs direct graphs is adding edges
// BFS and DFS logic remains the same

// for depth first traversal we use a stack
// DEPTH FIRST USES STACKS BUT STAy going down! (jokes)
// Recursively print elements depth first
func (d UndirectedGraph) PrintDepthFirstRecursion(s string, v map[string]bool) {
	fmt.Println(s)
	v[s] = true
	// recurisvely stack the neighbors
	for _, n := range d[s] {
		if !v[n] {
			d.PrintDepthFirstRecursion(n, v)
		}
	}
}

// Iteratively print elements depth first
func (d UndirectedGraph) PrintDepthFirstIterative(s string) {
	v := make(map[string]bool)
	stack := []string{s}

	// loop while we have items left in the stack
	for len(stack) > 0 {
		// Stacks are LIFO so use end
		i := len(stack) - 1
		e := stack[i]
		// remove last item from stack
		stack = stack[:i]
		// if we haven't visited e, print & add neighbors
		if !v[e] {
			// print & mark visited = true
			fmt.Println(e)
			v[e] = true
			// add neighbors
			for _, n := range d[e] {
				if !v[n] {
					stack = append(stack, n)
				}
			}
		}
	}
}

// for breadth first traversal we use a queue
// BREADTH FIRST USES QUEUES BUT QUE bread?? (jokes)
func (d UndirectedGraph) PrintBreadthFirst(s string) {
	v := make(map[string]bool)
	queue := []string{s}
	// loop while we have elements in our queue
	for len(queue) > 0 {
		// Queues are FIFO so use index 0
		e := queue[0]
		// remove element from queue
		queue = queue[1:]
		// if we haven't visited e, print & add neighbors
		if !v[e] {
			fmt.Println(e)
			v[e] = true
			// add neighbors to end of queue if not visited
			for _, n := range d[e] {
				if !v[n] {
					queue = append(queue, n)
				}
			}
		}
	}
}

func main() {
	ug := make(UndirectedGraph)

	// Representing friendships:
	//    Alice -- Bob
	//      |    /   |
	//     Carol  -- David
	addEdge(ug, "Alice", "Bob")
	addEdge(ug, "Alice", "Carol")
	addEdge(ug, "Bob", "Carol")
	addEdge(ug, "Bob", "David")
	addEdge(ug, "Carol", "David")

	// they actually won't match because recursion reverses
	// the output, but if we want to match we could
	// instead add neighbors in reverse in our iterative approach
	fmt.Println("Depth first recursion:")
	// make a map to know if we visited the nodes since we use recursion
	dFirstVisited := make(map[string]bool)
	ug.PrintDepthFirstRecursion("Alice", dFirstVisited)
	// iterative will make this map itself
	fmt.Println("Depth first iterative")
	ug.PrintDepthFirstIterative("Alice")
	fmt.Println("Breadth first:")
	ug.PrintBreadthFirst("Alice")
}
