package graph

import (
	"errors"
	"math"
	"sort"
)

type Edge struct {
	From, To int
	Weight   float64
}

func (e *Edge) Equals(e2 *Edge) bool {
	return e.From == e2.From && e.To == e2.To && e.Weight == e2.Weight
}

type WeightTuple struct {
	To     int
	Weight float64
}

// mapping from node index to list of tuples (node, weight)
type AdjList map[int][]WeightTuple

type Graph struct {
	AdjacencyList AdjList
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]WeightTuple)}
}

func FromEdgeList(edges []Edge) *Graph {
	g := NewGraph()
	for _, e := range edges {
		g.AddEdge(e.From, e.To, e.Weight)
	}
	return g
}

// adj is the weighted adjacency matrix
func FromAdjMat(adjMat []float64) (*Graph, error) {
	n := int(math.Sqrt(float64(len(adjMat))))
	// make sure that the adjacency matrix is square
	if n*n != len(adjMat) {
		return nil, errors.New("incorrect number of edges")
	}

	adjList := make(map[int][]WeightTuple, n)
	for i := 0; i < n; i++ {
		adjList[i] = []WeightTuple{}
		for j := 0; j < n; j++ {
			if adjMat[i*n+j] != 0 {
				tuple := WeightTuple{j, adjMat[i*n+j]}
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
	g.AdjacencyList[val] = []WeightTuple{}
}

func (g *Graph) HasNode(val int) bool {
	_, ok := g.AdjacencyList[val]
	return ok
}

func (g *Graph) AddEdge(from, to int, weight float64) {
	if g.Edge(from, to) != nil {
		return
	}
	if !g.HasNode(from) {
		g.AddNode(from)
	}
	if !g.HasNode(to) {
		g.AddNode(to)
	}
	g.AdjacencyList[from] = append(g.AdjacencyList[from], WeightTuple{to, weight})
}

func (g *Graph) Edge(from, to int) *Edge {
	adj := g.AdjacencyList[from]
	for i, e := range adj {
		if e.To == to {
			return &Edge{from, to, adj[i].Weight}
		}
	}
	return nil
}

func (g *Graph) Edges() []Edge {
	edges := make([]Edge, 0)
	for i, adj := range g.AdjacencyList {
		for _, e := range adj {
			edges = append(edges, Edge{i, e.To, e.Weight})
		}
	}
	return edges
}

func (g *Graph) Nodes() []int {
	nodes := []int{}
	for node := range g.AdjacencyList {
		nodes = append(nodes, node)
	}
	sort.Ints(nodes)
	return nodes
}

func (g *Graph) NumNodes() int {
	return len(g.AdjacencyList)
}

func (g *Graph) NumEdges() int {
	return len(g.Edges())
}

func (g *Graph) AdjEdges(i int) []Edge {
	edges := make([]Edge, 0)
	for _, e := range g.AdjacencyList[i] {
		edges = append(edges, Edge{i, e.To, e.Weight})
	}
	return edges
}

func (g *Graph) HasNegativeEdges() bool {
	for _, adj := range g.AdjacencyList {
		for _, e := range adj {
			if e.Weight < 0 {
				return true
			}
		}
	}
	return false
}
