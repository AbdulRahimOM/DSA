package graph

import (
	"fmt"
)

type Graph struct {
	adjacency map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacency: make(map[int][]int),
	}
}

func (graph *Graph) AddVertex(value int) error {
	if _, ok := graph.adjacency[value]; ok {
		return fmt.Errorf("vertex already exists")
	}
	graph.adjacency[value] = []int{}
	return nil
}

func (graph *Graph) AddEdge(from, to int) error {
	if _, ok := graph.adjacency[to]; !ok {
		return fmt.Errorf("to vertex doesn't exists")
	}
	if edges, ok := graph.adjacency[from]; !ok {
		return fmt.Errorf("from vertex doesn't exists")
	} else {
		if contains(edges, to) {
			return fmt.Errorf("edge already exist")
		} else {
			graph.adjacency[from] = append(edges, to)
		}
	}
	return nil
}
func contains(slice []int, key int) bool {
	for i := range slice {
		if slice[i] == key {
			return true
		}
	}
	return false
}

func (graph *Graph) PrintGraph() {
	for i,v:=range graph.adjacency{
		fmt.Printf("Vertex %d: %v\n",i,v)
	}
}