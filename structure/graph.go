package structure

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Node[T Number] struct {
	Val T // node value
}

type Edge[T Number] struct {
	From, To int // node indices
	Weight   T   // edge weight
}

type Graph[T Number] struct {
	Nodes []*Node[T]
	Edges []*Edge[T]
}

// n is the number of nodes in the graph
// adj is the weighted adjacency matrix
// node vals are the node values
func NewGraph[T Number](n int, adjList []T, nodeVal []T) (*Graph[T], error) {
	if n*n != len(adjList) {
		return nil, errors.New("incorrect number of edges")
	}

	nodes := make([]*Node[T], n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node[T]{nodeVal[i]}
	}

	edges := make([]*Edge[T], 0)
	for k := 0; k < n*n; k++ {
		i, j := k/n, k%n
		if adjList[k] != 0 {
			edges = append(edges, &Edge[T]{i, j, adjList[k]})
		}
	}

	g := &Graph[T]{nodes, edges}
	return g, nil
}

func (g *Graph[T]) AddNode(val T) {
	g.Nodes = append(g.Nodes, &Node[T]{val})
}

func (g *Graph[T]) AddEdge(from, to int) {
	g.AddWeightedEdge(from, to, 1)
}

func (g *Graph[T]) AddWeightedEdge(from, to int, weight T) {
	g.Edges = append(g.Edges, &Edge[T]{from, to, weight})
}
