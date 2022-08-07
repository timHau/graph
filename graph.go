package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

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

// n is the number of nodes in the graph
// adj is the weighted adjacency matrix
// node vals are the node values
func NewGraph[T any, N Number](n int, adjList []N, nodeVal []T) (*Graph[T, N], error) {
	// make sure that the adjacency matrix is square
	if n*n != len(adjList) {
		return nil, errors.New("incorrect number of edges")
	}
	// make sure that the node values are the same length as the number of nodes
	if n != len(nodeVal) {
		return nil, errors.New("incorrect number of node values")
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

// Breadth First Search
//
// Traverse the graph while exploring all nodes at the same level
// only explore nodes that have not been visited
//
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph[T, N]) BreadthFirstSearch(start int, fn func(*Node[T], int)) {
	visited := make([]bool, len(g.Nodes))
	g.BreadthFirstStep(start, visited, fn)

	// If the graph is not connected, we will start exploring the remaining graph components
	for i := 0; i < len(g.Nodes); i++ {
		if !visited[i] {
			g.BreadthFirstStep(i, visited, fn)
		}
	}
}

// Breadth First Step
//
// Perfomes a breadth first search step starting at node start
func (g *Graph[T, N]) BreadthFirstStep(start int, visited []bool, fn func(*Node[T], int)) {
	queue := make([]int, 0)
	queue = append(queue, start)
	visited[start] = true
	fn(g.Nodes[start], start)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, edge := range g.Edges {
			if edge.From == curr && !visited[edge.To] {
				queue = append(queue, edge.To)
				visited[edge.To] = true
				fn(g.Nodes[edge.To], edge.To)
			}
		}
	}
}

// Depth First Search
//
// # Traverse the graph while following the edges in a depth-first manner
//
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph[T, N]) DepthFirstSearch(start int, fn func(*Node[T], int)) {
	visited := make([]bool, len(g.Nodes))
	visited[start] = true

}
