package gograph

import (
	"fmt"
	"math"
)

func (g *Graph) BellmannFord(start int) {
	numNodes := g.NumNodes()
	dist := make([]float64, numNodes)
	parent := make([]int, numNodes)
	for i := 1; i < g.NumNodes(); i++ {
		dist[i] = math.Inf(1)
	}
	dist[start] = 0
	parent[start] = -1
	for i := 0; i < g.NumNodes()-1; i++ {
		for _, e := range g.AdjEdges(start) {
			if dist[e.To] > dist[start]+e.Weight {
				dist[e.To] = dist[start] + e.Weight
				parent[e.To] = start
			}
		}
	}
	for _, e := range g.AdjEdges(start) {
		if dist[e.To] > dist[start]+e.Weight {
			panic("negative cycle")
		}
	}

	fmt.Println("dist:", dist, "parent:", parent)
}
