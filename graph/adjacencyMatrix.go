package graph

import "fmt"

// Graph Space Complexity: O(n*n) Due to the use of 2D matrix

type MatrixGraph struct {
	adjMatrix [][]bool
	vertices  int
}

func (g *MatrixGraph) Init(vertices int) {
	g.vertices = vertices

	g.adjMatrix = make([][]bool, vertices)

	for i := 0; i < vertices; i++ {
		g.adjMatrix[i] = make([]bool, vertices)
	}

}

// AddEdge Time Complexity: O(1)
func (g *MatrixGraph) AddEdge(i, j int) {
	g.adjMatrix[i][j] = true
	g.adjMatrix[j][i] = true
}

// RemoveEdge Time Complexity: O(1)
func (g *MatrixGraph) RemoveEdge(i, j int) {
	g.adjMatrix[i][j] = false
	g.adjMatrix[j][i] = false
}

// AdjacentNodes Time Complexity: O(n)
func (g *MatrixGraph) AdjacentNodes(i int) {
	for vertex, isConnected := range g.adjMatrix[i] {
		if isConnected {
			fmt.Printf("%d\t", vertex)
		}
	}
}

// CheckIfConnected Time Complexity: O(1)
func (g *MatrixGraph) CheckIfConnected(i, j int) bool {
	if g.adjMatrix[i][j] {
		fmt.Printf("\n%d and %d are connected\n", i, j)
	} else {
		fmt.Printf("\n%d and %d are not connected\n", i, j)
	}

	return g.adjMatrix[i][j]
}

func (g *MatrixGraph) Display() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%d : ", i)
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%t ", g.adjMatrix[i][j])
		}

		fmt.Println()
	}
}
