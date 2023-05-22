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
