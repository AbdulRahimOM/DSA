package test

import (
	"DSA/DataStructures/graph"
	"fmt"
)

func TestGraph() {
	// Create a new graph
	graph := graph.NewGraph()
	if err:=graph.AddVertex(0);err!=nil{
		fmt.Println("error on adding vertex:",err)
	}
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)

	// Add edges to the graph
	if err:=graph.AddEdge(0, 1);err!=nil{
		fmt.Println("error on adding vertex:",err)
	}
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	// Print the graph
	fmt.Println("Graph Representation:")
	graph.PrintGraph()
}
