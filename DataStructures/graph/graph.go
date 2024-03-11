package graph

import (
	"fmt"
)

type Graph map[int][]int

func (graph *Graph) BFS(start int) {

	queue := []int{start}
	visited := map[int]bool{start: true}

	for len(queue) > 0 {
		fmt.Println("Visited :", queue[0])
		for _, neighbor := range (*graph)[queue[0]] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		queue = queue[1:]
	}
}

func (graph *Graph) DFS(node int){
	graph.dfs(node,make(map[int]bool))
}

func (graph *Graph) dfs(node int, visited map[int]bool) {
	visited[node]=true
	fmt.Println("Visited node:", node)
	for _,relative:=range (*graph)[node]{
		if !visited[relative]{
			graph.dfs(relative,visited)
		}
	}
}
func NewGraph() *Graph {
	graph := make(Graph)
	return &graph
}

func (graph *Graph) AddVertex(value int) error {
	if _, ok := (*graph)[value]; ok {
		return fmt.Errorf("vertex already exists")
	}
	(*graph)[value] = []int{}
	return nil
}

func (graph *Graph) AddEdge(from, to int) error {
	if _, ok := (*graph)[to]; !ok {
		return fmt.Errorf("to vertex doesn't exists")
	}
	if edges, ok := (*graph)[from]; !ok {
		return fmt.Errorf("from vertex doesn't exists")
	} else {
		if contains(edges, to) {
			return fmt.Errorf("edge already exist")
		} else {
			(*graph)[from] = append(edges, to)
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
	for i, v := range *graph {
		fmt.Printf("Vertex %d: %v\n", i, v)
	}
	fmt.Println()
}
