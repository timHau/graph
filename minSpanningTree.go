package gograph

import (
	"math"
)

func minKey(keys []float64, mstSet []bool) int {
	min := math.Inf(1)
	minIndex := 0
	for i, v := range keys {
		if v < min && !mstSet[i] {
			min = v
			minIndex = i
		}
	}
	return minIndex
}

func (g *Graph) Prim() *Graph {
	numNodes := g.NumNodes()
	mstSet := make([]bool, numNodes)     // Set of nodes in the MST
	mstKeys := make([]float64, numNodes) // dist values for each node
	for i := 1; i < g.NumNodes(); i++ {
		mstKeys[i] = math.Inf(1)
	}

	curr := 0
	mstKeys[curr] = 0
	res := NewGraph()

	for i := 0; i < g.NumNodes()-1; i++ {
		u := minKey(mstKeys, mstSet)
		mstSet[u] = true
		for _, e := range g.AdjEdges(u) {
			if !mstSet[e.To] && mstKeys[e.To] > e.Weight {
				mstKeys[e.To] = e.Weight
				res.AddWeightedEdge(u, e.To, e.Weight)
			}
		}
	}

	return res
}
