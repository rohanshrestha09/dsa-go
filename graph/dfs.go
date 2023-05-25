package graph

func (g *Graph[T]) DFS(start T) []T {

	visited := make(map[T]bool)

	var path []T

	var dfs func(vertex T)

	dfs = func(vertex T) {

		if visited[vertex] {
			return
		}

		path = append(path, vertex)

		visited[vertex] = true

		temp := g.list[vertex]

		for temp != nil {
			dfs(temp.vertex)

			temp = temp.link
		}
	}

	dfs(start)

	return path
}
