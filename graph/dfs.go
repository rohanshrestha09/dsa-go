package graph

import "fmt"

func (g *Graph) DFS(initialVertex int, visited []bool) {

	if visited[initialVertex] {
		return
	}

	fmt.Print(initialVertex)

	visited[initialVertex] = true

	temp := g.adjList[initialVertex]

	for temp != nil {
		g.DFS(temp.vertex, visited)

		temp = temp.link
	}
}

func dfs() {

	g := new(Graph)

	g.Init(5)

	g.AddEdge(0, 3)

	g.AddEdge(0, 2)

	g.AddEdge(0, 1)

	g.AddEdge(2, 4)

	visited := make([]bool, g.vertices)

	g.DFS(0, visited)
}
