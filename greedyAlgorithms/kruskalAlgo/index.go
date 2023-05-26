package kruskalAlgo

import (
	"fmt"
)

type Graph[T string | int] struct {
	list []*Path[T]
}

type Path[T string | int] struct {
	node   []T
	weight uint
}

func (g *Graph[T]) AddEdge(s, d T, w uint) {
	if s == d {
		return
	}

	g.list = append(g.list, &Path[T]{node: []T{s, d}, weight: w})
}

func (g *Graph[T]) RemoveEdge(s, d T) {
	for i, v := range g.list {
		if intersection := g.FindIntersection(v.node, []T{s, d}); len(intersection) == 2 {
			g.list = append(g.list[:i], g.list[i+1:]...)
		}
	}
}

func (g *Graph[T]) FindUnion(A, B []T) []T {
	union := make([]T, 0)

	elements := make(map[T]bool)

	for _, v := range A {
		elements[v] = true
		union = append(union, v)
	}

	for _, v := range B {
		if !elements[v] {
			elements[v] = true
			union = append(union, v)
		}
	}

	return union
}

func (g *Graph[T]) FindIntersection(A, B []T) []T {
	intersection := make([]T, 0)

	elements := make(map[T]bool)

	for _, v := range A {
		elements[v] = true
	}

	for _, v := range B {
		if elements[v] {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func (g *Graph[T]) Sort() {
	for i := 1; i < len(g.list); i++ {
		k := i

		for k-1 >= 0 && g.list[k].weight < g.list[k-1].weight {
			temp := g.list[k-1]

			g.list[k-1] = g.list[k]

			g.list[k] = temp

			k--
		}
	}
}

func (g *Graph[T]) KruskalAlgo() []Path[T] {

	var visited []T

	var list []Path[T]

	g.Sort()

	if len(g.list) == 0 {
		return nil
	}

	for _, v := range g.list {
		intersection := g.FindIntersection(visited, v.node)

		if len(intersection) != 2 {
			list = append(list, *v)
		}

		visited = g.FindUnion(visited, v.node)
	}

	return list

}

func Test() {
	g := new(Graph[string])

	g.AddEdge("A", "B", 4)

	g.AddEdge("A", "C", 4)

	g.AddEdge("B", "C", 2)

	g.AddEdge("C", "D", 3)

	g.AddEdge("C", "E", 2)

	g.AddEdge("C", "F", 4)

	g.AddEdge("D", "F", 3)

	g.AddEdge("E", "F", 3)

	fmt.Println(g.KruskalAlgo())

}
