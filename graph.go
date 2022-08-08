package gograph

import (
	"errors"
	"math"
)

type Edge struct {
	From, To int
	Weight   float64
}

type AdjListTuple struct {
	to     int
	weight float64
}

// mapping from node index to list of tuples (node, weight)
type AdjList map[int][]AdjListTuple

type Graph struct {
	AdjacencyList AdjList
}

func New() *Graph {
	return &Graph{make(map[int][]AdjListTuple)}
}

// adj is the weighted adjacency matrix
func FromAdjMat(adjMat []float64) (*Graph, error) {
	n := int(math.Sqrt(float64(len(adjMat))))
	// make sure that the adjacency matrix is square
	if n*n != len(adjMat) {
		return nil, errors.New("incorrect number of edges")
	}

	adjList := make(map[int][]AdjListTuple)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if adjMat[i*n+j] != 0 {
				tuple := AdjListTuple{j, adjMat[i*n+j]}
				adjList[i] = append(adjList[i], tuple)
			}
		}
	}

	return &Graph{adjList}, nil
}

func FromAdjList(adjList AdjList) *Graph {
	return &Graph{adjList}
}

func (g *Graph) AddNode(val int) {
	g.AdjacencyList[val] = []AdjListTuple{}
}

func (g *Graph) AddEdge(from, to int) {
	g.AddWeightedEdge(from, to, 1)
}

func (g *Graph) AddWeightedEdge(from, to int, weight float64) {
	if g.Edge(from, to) != nil {
		return
	}
	g.AdjacencyList[from] = append(g.AdjacencyList[from], AdjListTuple{to, weight})
}

func (g *Graph) Edge(from, to int) *Edge {
	adj := g.AdjacencyList[from]
	for i, e := range adj {
		if e.to == to {
			return &Edge{from, to, adj[i].weight}
		}
	}
	return nil
}

func (g *Graph) Edges() []Edge {
	edges := make([]Edge, 0)
	for i, adj := range g.AdjacencyList {
		for _, e := range adj {
			edges = append(edges, Edge{i, e.to, e.weight})
		}
	}
	return edges
}

func (g *Graph) Nodes() []int {
	nodes := []int{}
	for node := range g.AdjacencyList {
		nodes = append(nodes, node)
	}
	return nodes
}

func (g *Graph) NumNodes() int {
	return len(g.AdjacencyList)
}

func (g *Graph) NumEdges() int {
	return len(g.Edges())
}

func (g *Graph) Neighbors(i int) []AdjListTuple {
	return g.AdjacencyList[i]
}
