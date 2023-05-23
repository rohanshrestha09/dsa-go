package graph

import "fmt"

func (g *Graph[T]) DFS(initialVertex T, visited map[T]bool) {

	if visited[initialVertex] {
		return
	}

	fmt.Printf("%v\t", initialVertex)

	visited[initialVertex] = true

	temp := g.adjList[initialVertex]

	for temp != nil {
		g.DFS(temp.vertex, visited)

		temp = temp.link
	}
}
