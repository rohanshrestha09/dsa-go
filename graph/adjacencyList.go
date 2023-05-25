package graph

import "fmt"

type GraphType int

const (
	DIRECTED GraphType = iota
	UNDIRECTED
)

type Node[T string | int] struct {
	vertex T
	link   *Node[T]
}

type Graph[T string | int] struct {
	list      map[T]*Node[T]
	vertices  int
	graphType GraphType
}

func (g *Graph[T]) Init(vertices int, graphType GraphType) {
	g.vertices = vertices

	g.list = make(map[T]*Node[T], vertices)

	g.graphType = graphType
}

func (g *Graph[T]) AddEdge(s, d T) {

	newNode := &Node[T]{d, g.list[s]}

	g.list[s] = newNode

	if g.graphType == UNDIRECTED {
		newNode := &Node[T]{s, g.list[d]}

		g.list[d] = newNode
	}
}

func (g *Graph[T]) RemoveEdge(s, d T) {
	if g.list[s] == nil {
		return
	}

	if g.list[s].vertex == d {
		g.list[s] = g.list[s].link

		if g.graphType == DIRECTED {
			return
		}

		g.RemoveEdge(d, s)
	}

	sourceVertex := g.list[s]

	for sourceVertex.link != nil {
		if sourceVertex.link.vertex == d {
			sourceVertex.link = sourceVertex.link.link

			if g.graphType == DIRECTED {
				return
			}

			g.RemoveEdge(d, s)
		}

		sourceVertex = sourceVertex.link
	}
}

func (g *Graph[T]) AdjacentNodes(s T) {
	sourceVertex := g.list[s]

	fmt.Printf("Adjacent nodes of %v: ", s)
	for sourceVertex != nil {
		fmt.Printf("%v\t", sourceVertex.vertex)

		sourceVertex = sourceVertex.link
	}
}

func (g *Graph[T]) Display() {
	for k := range g.list {
		fmt.Printf("\n%v: ", k)
		for sourceVertex := g.list[k]; sourceVertex != nil; sourceVertex = sourceVertex.link {
			fmt.Printf("%v\t", sourceVertex.vertex)
		}
	}
}
