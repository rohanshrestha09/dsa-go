//go:build exclude

package graph

import "fmt"

// Graph Space Complexity: O(n) Due to the use of 2D matrix

type Graph struct {
	adjMatrix [][]bool
	vertices  int
}

func (graph *Graph) Init(vertices int) {
	graph.vertices = vertices

	graph.adjMatrix = make([][]bool, vertices)

	for i := 0; i < vertices; i++ {
		graph.adjMatrix[i] = make([]bool, vertices)
	}

}

// AddEdge Time Complexity: O(1)
func (graph *Graph) AddEdge(i, j int) {
	graph.adjMatrix[i][j] = true
	graph.adjMatrix[j][i] = true
}

// RemoveEdge Time Complexity: O(1)
func (graph *Graph) RemoveEdge(i, j int) {
	graph.adjMatrix[i][j] = false
	graph.adjMatrix[j][i] = false
}

// AdjacentNodes Time Complexity: O(n)
func (graph *Graph) AdjacentNodes(i int) {
	for vertex, isConnected := range graph.adjMatrix[i] {
		if isConnected {
			fmt.Printf("%d\t", vertex)
		}
	}
}

// CheckIfConnected Time Complexity: O(1)
func (graph *Graph) CheckIfConnected(i, j int) bool {
	if graph.adjMatrix[i][j] {
		fmt.Printf("\n%d and %d are connected\n", i, j)
	} else {
		fmt.Printf("\n%d and %d are not connected\n", i, j)
	}

	return graph.adjMatrix[i][j]
}

func (graph *Graph) Display() {
	for i := 0; i < graph.vertices; i++ {
		fmt.Printf("%d : ", i)
		for j := 0; j < graph.vertices; j++ {
			fmt.Printf("%t ", graph.adjMatrix[i][j])
		}

		fmt.Println()
	}
}

func adjacencyMatrix() {
	graph := new(Graph)

	graph.Init(4)

	graph.AddEdge(1, 2)

	graph.AddEdge(0, 3)

	graph.AddEdge(1, 3)

	graph.AddEdge(0, 2)

	graph.RemoveEdge(2, 1)

	graph.Display()

	graph.CheckIfConnected(0, 3)

	graph.AdjacentNodes(0)
}
