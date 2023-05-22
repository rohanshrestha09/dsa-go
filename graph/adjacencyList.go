package graph

import "fmt"

type Node struct {
	vertex int
	link   *Node
}

type Graph struct {
	adjList  []*Node
	vertices int
}

func (g *Graph) Init(vertices int) {
	g.vertices = vertices

	g.adjList = make([]*Node, vertices)
}

func (g *Graph) AddEdge(s, d int) {
	newNode := &Node{d, nil}

	newNode.link = g.adjList[s]

	g.adjList[s] = newNode

	if g.adjList[d] == nil {
		g.AddEdge(d, s)
	}
}

func (g *Graph) RemoveEdge(s, d int) {
	if g.adjList[s] == nil {
		return
	}

	if g.adjList[s].vertex == d {
		g.adjList[s] = g.adjList[s].link
		g.RemoveEdge(d, s)
		return
	}

	sourceVertex := g.adjList[s]

	for sourceVertex.link != nil {
		if sourceVertex.link.vertex == d {
			sourceVertex.link = sourceVertex.link.link

			g.RemoveEdge(d, s)

			return
		}
		sourceVertex = sourceVertex.link
	}
}

func (g *Graph) AdjacentNodes(s byte) {
	sourceVertex := g.adjList[s]

	fmt.Printf("Adjacent nodes of %d: ", s)
	for sourceVertex != nil {
		fmt.Printf("%v\t", sourceVertex.vertex)

		sourceVertex = sourceVertex.link
	}
}

func (g *Graph) Display() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("\n%d: ", i)
		for sourceVertex := g.adjList[i]; sourceVertex != nil; sourceVertex = sourceVertex.link {
			fmt.Printf("%v\t", sourceVertex.vertex)
		}
	}
}
