package primsAlgo

import (
	"fmt"
	"math"
)

type Node[T string | int] struct {
	src    T
	weight uint
	dest   *Node[T]
}

type Graph[T string | int] struct {
	list     map[T]*Node[T]
	vertices int
}

type Path[T string | int] struct {
	node   []T
	weight uint
}

func (g *Graph[T]) Init(vertices int) {
	g.vertices = vertices

	g.list = make(map[T]*Node[T], vertices)
}

func (g *Graph[T]) AddEdge(s, d T, w uint) {

	if s == d {
		return
	}

	newNode := &Node[T]{d, w, g.list[s]}

	g.list[s] = newNode

	newNode = &Node[T]{s, w, g.list[d]}

	g.list[d] = newNode
}

func (g *Graph[T]) RemoveEdge(s, d T) {

	if g.list[s] == nil {
		return
	}

	if g.list[s].src == d {
		g.list[s] = g.list[s].dest

		g.RemoveEdge(d, s)
	}

	source := g.list[s]

	for source.dest != nil {
		if source.dest.src == d {
			source.dest = source.dest.dest

			g.RemoveEdge(d, s)
		}

		source = source.dest
	}
}

func (g *Graph[T]) MinAdjNode(src T, unvisited map[T]bool) *Node[T] {
	var minCost uint = math.MaxUint32

	var minNode *Node[T]

	current := g.list[src]

	for current != nil {
		if unvisited == nil || unvisited[current.src] {
			if current.weight < minCost {
				minCost, minNode = current.weight, current
			}
		}
		current = current.dest
	}

	return minNode
}

func (g *Graph[T]) PrimsAlgo(vertex T) []Path[T] {
	unvisited := make(map[T]bool)

	for k := range g.list {
		unvisited[k] = true
	}

	var path []Path[T]

	var primsAlgo func(v T)

	primsAlgo = func(v T) {

		minAdjNode := g.MinAdjNode(v, unvisited)

		delete(unvisited, v)

		if minAdjNode == nil {
			for k := range unvisited {

				minAdjNode := g.MinAdjNode(k, nil)

				unvisited[minAdjNode.src] = true

				primsAlgo(k)

				return
			}

			return
		}

		path = append(path, Path[T]{
			node:   []T{v, minAdjNode.src},
			weight: minAdjNode.weight,
		})

		if len(unvisited) != 0 {
			primsAlgo(minAdjNode.src)
		}
	}

	primsAlgo(vertex)

	return path
}

func Test() {
	g := new(Graph[string])

	g.Init(6)

	g.AddEdge("A", "B", 4)

	g.AddEdge("A", "C", 4)

	g.AddEdge("B", "C", 2)

	g.AddEdge("C", "D", 3)

	g.AddEdge("C", "E", 2)

	g.AddEdge("C", "F", 4)

	g.AddEdge("D", "F", 3)

	g.AddEdge("E", "F", 3)

	fmt.Println(g.PrimsAlgo("D"))
}
