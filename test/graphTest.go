package test

import (
	"DSA/DataStructures/graph"
	"fmt"
)

func TestGraph1() {
	// Example graph represented as an adjacency list
	graph := graph.Graph{
		0: {1, 2},
		1: {3},
		2: {4, 6},
		3: {5},
		4: {},
		5: {},
		6: {},
	}

	// Print the graph
	fmt.Println("Graph Representation:")
	graph.PrintGraph()

	// Perform DFS starting from node 0
	fmt.Println("DFS:")
	graph.LookUpByDFS(0)

	// Perform LookUpByBFS starting from node 0
	fmt.Println("\nBFS:")
	graph.LookUpByBFS(0)
}

func TestGraph2() {

	// Create a new graph
	graph := graph.NewGraph()
	if err := graph.AddVertex(0); err != nil {
		fmt.Println("error on adding vertex:", err)
	}
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	// Add edges to the graph
	if err := graph.AddEdge(0, 1); err != nil {
		fmt.Println("error on adding vertex:", err)
	}
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	// Print the graph
	fmt.Println("Graph Representation:")
	graph.PrintGraph()

	// Perform DFS starting from node 0
	fmt.Println("DFS:")
	graph.LookUpByDFS(0)

	// Perform LookUpByBFS starting from node 0
	fmt.Println("\nBFS:")
	graph.LookUpByBFS(0)
}
