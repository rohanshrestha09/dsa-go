package graph

import (
	"DSA-Go/queue"
)

func (g *Graph[T]) BFS(start T) (map[T]T, []T) {

	var path []T

	queue := new(queue.ListQueue[T])

	visited := make(map[T]T, g.vertices)

	queue.Enqueue(start)

	visited[start] = g.list[start].vertex

	for queue.Length != 0 {
		current := queue.Dequeue()

		path = append(path, current)

		temp := g.list[current]

		for temp != nil {
			adjVertex := temp.vertex

			if _, ok := visited[adjVertex]; !ok {
				queue.Enqueue(adjVertex)

				visited[adjVertex] = current
			}

			temp = temp.link
		}
	}

	return visited, path
}

func (g *Graph[T]) ShortestPath(source, target T) []T {
	var path []T

	visited, _ := g.BFS(source)

	if _, ok := visited[target]; !ok {
		return nil
	}

	current := target

	for current != source {
		path = append(path, current)

		current = visited[current]
	}

	path = append(path, source)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
