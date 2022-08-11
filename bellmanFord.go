package graph

import (
	"errors"
	"math"
)

func (g *Graph) BellmanFord(start int) ([]float64, error) {
	numNodes := g.NumNodes()
	dist := make([]float64, numNodes)
	for i := 1; i < g.NumNodes(); i++ {
		dist[i] = math.Inf(1)
	}
	dist[start] = 0

	for i := 0; i < g.NumNodes()-1; i++ {
		for _, e := range g.Edges() {
			if dist[e.From] != math.Inf(1) && dist[e.To] > dist[e.From]+e.Weight {
				dist[e.To] = dist[e.From] + e.Weight
			}
		}
	}

	for _, e := range g.Edges() {
		if dist[e.From] != math.Inf(1) && dist[e.To] > dist[e.From]+e.Weight {
			return nil, errors.New("Graph contains negative cycle")
		}
	}

	return dist, nil
}
