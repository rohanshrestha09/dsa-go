package graph

import (
	"DSA-Go/queue"
	"fmt"
)

func (g *Graph[T]) BFS(initialVertex T) {

	queue := new(queue.ListQueue[T])

	visited := make(map[T]bool, g.vertices)

	queue.Enqueue(initialVertex)

	visited[initialVertex] = true

	for queue.Length != 0 {
		currentVertex := queue.Dequeue()

		fmt.Printf("%v\t", currentVertex)

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
