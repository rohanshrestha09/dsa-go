package graph

import "fmt"

func Run() {

	g := new(Graph)

	g.Init(5)

	g.AddEdge(0, 3)

	g.AddEdge(0, 2)

	g.AddEdge(0, 1)

	g.AddEdge(2, 4)

	g.BFS(0)

	fmt.Println()

	visited := make([]bool, g.vertices)

	g.DFS(0, visited)
}
