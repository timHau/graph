package graph

import "math"

// Floyd Warshall algorithm for finding all-pairs shortest paths in a graph.
//
// returns a matrix, given as a flat array of shortest distances between all nodes in the graph
//
// Time Complexity: O(V^3)
func (g *Graph) FloydWarshall() []float64 {
	numNodes := g.NumNodes()
	dist := make([]float64, numNodes*numNodes)
	for i := 0; i < numNodes*numNodes; i++ {
		dist[i] = math.Inf(1)
	}
	ind := index(numNodes)

	for _, e := range g.Edges() {
		dist[ind(e.From, e.To)] = e.Weight
	}

	for k := 0; k < numNodes; k++ {
		dist[ind(k, k)] = 0 // distance from a node to itself is 0
	}

	for k := 0; k < numNodes; k++ {
		for i := 0; i < numNodes; i++ {
			for j := 0; j < numNodes; j++ {
				if dist[ind(i, k)]+dist[ind(k, j)] < dist[ind(i, j)] {
					dist[ind(i, j)] = dist[ind(i, k)] + dist[ind(k, j)]
				}
			}
		}
	}

	return dist
}

func index(n int) func(int, int) int {
	return func(i, j int) int {
		return n*i + j
	}
}
