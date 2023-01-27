package graph

import "fmt"

type Node struct {
	vertex byte
	link   *Node
}

type Graph struct {
	adjList  []*Node
	vertices int
}

func (graph *Graph) Init(vertices int) {
	graph.vertices = vertices

	graph.adjList = make([]*Node, vertices)
}

func (graph *Graph) AddEdge(s, d byte) {
	newNode := &Node{d, nil}

	newNode.link = graph.adjList[s]

	graph.adjList[s] = newNode

	if graph.adjList[d] == nil {
		graph.AddEdge(d, s)
	}
}

func (graph *Graph) RemoveEdge(s, d byte) {
	if graph.adjList[s] == nil {
		return
	}

	if graph.adjList[s].vertex == d {
		graph.adjList[s] = graph.adjList[s].link
		graph.RemoveEdge(d, s)
		return
	}

	sourceVertex := graph.adjList[s]

	for sourceVertex.link != nil {
		if sourceVertex.link.vertex == d {
			sourceVertex.link = sourceVertex.link.link

			graph.RemoveEdge(d, s)

			return
		}
		sourceVertex = sourceVertex.link
	}
}

func (graph *Graph) AdjacentNodes(s byte) {
	sourceVertex := graph.adjList[s]

	fmt.Printf("Adjacent nodes of %d: ", s)
	for sourceVertex != nil {
		fmt.Printf("%v\t", sourceVertex.vertex)

		sourceVertex = sourceVertex.link
	}
}

func (graph *Graph) Display() {
	for i := 0; i < graph.vertices; i++ {
		fmt.Printf("\n%d: ", i)
		for sourceVertex := graph.adjList[i]; sourceVertex != nil; sourceVertex = sourceVertex.link {
			fmt.Printf("%v\t", sourceVertex.vertex)
		}
	}
}

func adjacencyList() {
	graph := new(Graph)

	graph.Init(4)

	graph.AddEdge(0, 3)

	graph.AddEdge(0, 1)

	graph.AddEdge(1, 2)

	graph.AdjacentNodes(0)

	graph.RemoveEdge(3, 0)

	graph.Display()
}
