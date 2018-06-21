// Graph maps a string to a set of related strings, its successors in a directed graph
package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	fmt.Println(hasEdge("a", "b"))
	addEdge("a", "b")
	fmt.Println(hasEdge("a", "b"))
}
