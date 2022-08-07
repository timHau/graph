package gograph

import (
	"errors"
	"fmt"
)

type AdjListTuple struct {
	to     int
	weight float64
}

// mapping from node index to list of tuples (node, weight)
type AdjList map[int][]AdjListTuple

type Node[T any] struct {
	Val T // node value
}

type Edge struct {
	From, To int     // node indices
	Weight   float64 // edge weight
}

type Graph[T any] struct {
	Nodes []*Node[T]
	Edges []*Edge
}

// adj is the weighted adjacency matrix
// node vals are the node values
func NewGraph[T any](adjList []float64, nodeVal []T) (*Graph[T], error) {
	n := len(nodeVal)
	// make sure that the adjacency matrix is square
	if n*n != len(adjList) {
		return nil, errors.New("incorrect number of edges")
	}

	nodes := make([]*Node[T], n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node[T]{nodeVal[i]}
	}

	edges := make([]*Edge, 0)
	for k := 0; k < n*n; k++ {
		i, j := k/n, k%n
		if adjList[k] != 0 {
			edges = append(edges, &Edge{i, j, adjList[k]})
		}
	}

	return &Graph[T]{nodes, edges}, nil
}

func (g *Graph[T]) FromMap(adjList map[int][]int, nodeVal []T) (*Graph[T], error) {
	n := len(nodeVal)
	if n != len(adjList) {
		return nil, errors.New("incorrect number of nodes")
	}

	nodes := make([]*Node[T], n)
	edges := make([]*Edge, 0)
	for k, v := range adjList {
		nodes[k] = &Node[T]{nodeVal[k]}
		for _, j := range v {
			edges = append(edges, &Edge{k, j, 1})
		}
	}

	return &Graph[T]{nodes, edges}, nil
}

func (g *Graph[T]) AddNode(val T) {
	g.Nodes = append(g.Nodes, &Node[T]{val})
}

func (g *Graph[T]) AddEdge(from, to int) {
	g.AddWeightedEdge(from, to, 1)
}

func (g *Graph[T]) AddWeightedEdge(from, to int, weight float64) {
	g.Edges = append(g.Edges, &Edge{from, to, weight})
}

func (g *Graph[T]) NumNodes() int {
	return len(g.Nodes)
}

func (g *Graph[T]) NumEdges() int {
	return len(g.Edges)
}

func (g *Graph[T]) AsAdjList() AdjList {
	adjList := make(map[int][]AdjListTuple)
	for _, e := range g.Edges {
		adjList[e.From] = append(adjList[e.From], AdjListTuple{e.To, e.Weight})
	}
	return adjList
}

func (g *Graph[T]) Neighbors(i int) []AdjListTuple {
	return g.AsAdjList()[i]
}

func (g *Graph[T]) Edge(from, to int) *Edge {
	for _, e := range g.Edges {
		if e.From == from && e.To == to {
			return e
		}
	}
	return nil
}

func (g *Graph[T]) Show() {
	adjList := g.AsAdjList()
	for i, _ := range g.Nodes {
		println(fmt.Sprintf("%d: %v", i, adjList[i]))
	}
}
