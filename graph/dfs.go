package graph

import "fmt"

func (g *Graph) DFS(initialVertex int, visited []*Node) {

	if visited[initialVertex] != nil {
		return
	}

	fmt.Print(initialVertex)

	temp := g.adjList[initialVertex]

	visited[initialVertex] = temp

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

	visited := make([]*Node, 5)

	g.DFS(0, visited)
}
