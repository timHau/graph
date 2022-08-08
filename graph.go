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

type Node struct {
	ID int
}

type Edge struct {
	From, To int     // node indices
	Weight   float64 // edge weight
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// adj is the weighted adjacency matrix
// node vals are the node values
func NewGraph(adjList []float64, nodeVal []int) (*Graph, error) {
	n := len(nodeVal)
	// make sure that the adjacency matrix is square
	if n*n != len(adjList) {
		return nil, errors.New("incorrect number of edges")
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{nodeVal[i]}
	}

	edges := make([]*Edge, 0)
	for k := 0; k < n*n; k++ {
		i, j := k/n, k%n
		if adjList[k] != 0 {
			edges = append(edges, &Edge{i, j, adjList[k]})
		}
	}

	return &Graph{nodes, edges}, nil
}

func FromAdjList(adjList AdjList, nodeVal []int) (*Graph, error) {
	n := len(nodeVal)
	if n != len(adjList) {
		return nil, errors.New("incorrect number of nodes")
	}

	nodes := make([]*Node, n)
	edges := make([]*Edge, 0)
	for k, v := range adjList {
		nodes[k] = &Node{nodeVal[k]}
		for _, j := range v {
			edges = append(edges, &Edge{k, j.to, j.weight})
		}
	}

	return &Graph{nodes, edges}, nil
}

func (g *Graph) AddNode(val int) {
	g.Nodes = append(g.Nodes, &Node{val})
}

func (g *Graph) AddEdge(from, to int) {
	g.AddWeightedEdge(from, to, 1)
}

func (g *Graph) AddWeightedEdge(from, to int, weight float64) {
	g.Edges = append(g.Edges, &Edge{from, to, weight})
}

func (g *Graph) NumNodes() int {
	return len(g.Nodes)
}

func (g *Graph) NumEdges() int {
	return len(g.Edges)
}

func (g *Graph) AsAdjList() AdjList {
	adjList := make(map[int][]AdjListTuple)
	for _, e := range g.Edges {
		adjList[e.From] = append(adjList[e.From], AdjListTuple{e.To, e.Weight})
	}
	return adjList
}

func (g *Graph) Neighbors(i int) []AdjListTuple {
	return g.AsAdjList()[i]
}

func (g *Graph) Edge(from, to int) *Edge {
	for _, e := range g.Edges {
		if e.From == from && e.To == to {
			return e
		}
	}
	return nil
}

func (g *Graph) Show() {
	adjList := g.AsAdjList()
	for i, _ := range g.Nodes {
		println(fmt.Sprintf("%d: %v", i, adjList[i]))
	}
}
