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

	_, path := g.BFS("Rohan")

	fmt.Println(path)

	fmt.Println(g.ShortestPath("Rohan", "Kriti"))

	fmt.Println(g.DFS("Rohan"))

	g.RemoveEdge("Kriti", "Shrestha")

	g.RemoveEdge("Rohan", "Shrestha")

	g.Display()

	fmt.Println()

	g.AdjacentNodes("Rose")
}
