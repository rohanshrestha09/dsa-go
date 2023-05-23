package graph

import "fmt"

var data = map[string][]string{
	"Rohan":    {"Shrestha", "John"},
	"Shrestha": {"Rohan", "James", "Kriti"},
	"John":     {"Rohan", "Rose"},
	"James":    {"Shrestha", "Rose"},
	"Kriti":    {"Shrestha"},
	"Rose":     {"James", "John"},
}

func Run() {

	g := new(Graph[string])

	g.Init(6, DIRECTED)

	for k, v := range data {
		for _, v := range v {
			g.AddEdge(k, v)
		}
	}

	g.BFS("Rohan")

	fmt.Println()

	visited := make(map[string]bool, g.vertices)

	g.DFS("Rohan", visited)

	g.RemoveEdge("Kriti", "Shrestha")

	g.RemoveEdge("Rohan", "Shrestha")

	g.Display()

	fmt.Println()

	g.AdjacentNodes("Rose")
}
