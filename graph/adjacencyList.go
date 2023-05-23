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
	adjList   map[T]*Node[T]
	vertices  int
	graphType GraphType
}

func (g *Graph[T]) Init(vertices int, graphType GraphType) {
	g.vertices = vertices

	g.adjList = make(map[T]*Node[T], vertices)

	g.graphType = graphType
}

func (g *Graph[T]) AddEdge(s, d T) {

	newNode := &Node[T]{d, nil}

	newNode.link = g.adjList[s]

	g.adjList[s] = newNode

	if g.graphType == DIRECTED {
		return
	}

	if g.adjList[d] == nil {
		g.AddEdge(d, s)
	}
}

func (g *Graph[T]) RemoveEdge(s, d T) {
	if g.adjList[s] == nil {
		return
	}

	if g.adjList[s].vertex == d {
		g.adjList[s] = g.adjList[s].link

		if g.graphType == DIRECTED {
			return
		}

		g.RemoveEdge(d, s)
	}

	sourceVertex := g.adjList[s]

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
	sourceVertex := g.adjList[s]

	fmt.Printf("Adjacent nodes of %v: ", s)
	for sourceVertex != nil {
		fmt.Printf("%v\t", sourceVertex.vertex)

		sourceVertex = sourceVertex.link
	}
}

func (g *Graph[T]) Display() {
	for k := range g.adjList {
		fmt.Printf("\n%v: ", k)
		for sourceVertex := g.adjList[k]; sourceVertex != nil; sourceVertex = sourceVertex.link {
			fmt.Printf("%v\t", sourceVertex.vertex)
		}
	}
}
