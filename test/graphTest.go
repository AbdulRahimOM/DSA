package test

import (
	"DSA/DataStructures/graph"
	"fmt"
)

func TestGraph() {
	// Create a new graph
	graph := graph.NewGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	// graph.AddVertex(2)

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	// Print the graph
	fmt.Println("Graph Representation:")
	graph.PrintGraph()
}
