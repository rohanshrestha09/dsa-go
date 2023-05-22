package graph

import (
	"DSA-Go/queue"
	"fmt"
)

func (g *Graph) BFS(initialVertex int) {

	queue := new(queue.ListQueue[int])

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
