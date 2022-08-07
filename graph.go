package gograph

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type AdjList map[int][]int

type Node[T any] struct {
	Val T // node value
}

type Edge[T any] struct {
	From, To int // node indices
	Weight   T   // edge weight
}

type Graph[T any, N Number] struct {
	Nodes []*Node[T]
	Edges []*Edge[N]
}

// adj is the weighted adjacency matrix
// node vals are the node values
func NewGraph[T any, N Number](adjList []N, nodeVal []T) (*Graph[T, N], error) {
	n := len(nodeVal)
	// make sure that the adjacency matrix is square
	if n*n != len(adjList) {
		return nil, errors.New("incorrect number of edges")
	}

	nodes := make([]*Node[T], n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node[T]{nodeVal[i]}
	}

	edges := make([]*Edge[N], 0)
	for k := 0; k < n*n; k++ {
		i, j := k/n, k%n
		if adjList[k] != 0 {
			edges = append(edges, &Edge[N]{i, j, adjList[k]})
		}
	}

	g := &Graph[T, N]{nodes, edges}
	return g, nil
}

func (g *Graph[T, N]) AddNode(val T) {
	g.Nodes = append(g.Nodes, &Node[T]{val})
}

func (g *Graph[T, N]) AddEdge(from, to int) {
	g.AddWeightedEdge(from, to, 1)
}

func (g *Graph[T, N]) AddWeightedEdge(from, to int, weight N) {
	g.Edges = append(g.Edges, &Edge[N]{from, to, weight})
}

func (g *Graph[T, N]) NumNodes() int {
	return len(g.Nodes)
}

func (g *Graph[T, N]) NumEdges() int {
	return len(g.Edges)
}

func (g *Graph[T, N]) AsAdjList() AdjList {
	adjList := make(map[int][]int)
	for _, e := range g.Edges {
		adjList[e.From] = append(adjList[e.From], e.To)
	}
	return adjList
}

func (g *Graph[T, N]) Neighbors(i int) []int {
	return g.AsAdjList()[i]
}

func (g *Graph[T, N]) Show() {
	adjList := g.AsAdjList()
	for i, _ := range g.Nodes {
		println(fmt.Sprintf("%d: %v", i, adjList[i]))
	}
}
