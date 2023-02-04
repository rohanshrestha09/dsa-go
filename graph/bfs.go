package graph

import (
	queue "DSA-Go/queue"
	"fmt"
)

func (g *Graph) BFS(initialVertex int) {

	queue := new(queue.ListQueue)

	visited := make([]bool, g.vertices)

	queue.Enqueue(initialVertex)

	visited[initialVertex] = true

	for queue.Length != 0 {
		currentVertex := queue.Dequeue()

		fmt.Print(currentVertex)

		temp := g.adjList[currentVertex]

		for temp != nil {
			adjVertex := temp.vertex

			if !visited[adjVertex] {
				queue.Enqueue(adjVertex)

				visited[adjVertex] = true
			}

			temp = temp.link
		}

	}

}

func bfs() {

	g := new(Graph)

	g.Init(5)

	g.AddEdge(0, 3)

	g.AddEdge(0, 2)

	g.AddEdge(0, 1)

	g.AddEdge(2, 4)

	g.BFS(0)

}
