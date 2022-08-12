package graph

// Laplacian returns the Laplacian matrix of a graph as a flat array.
func (g *Graph) Laplacian() []float64 {
	numNodes := g.NumNodes()
	degrees := make([]int, numNodes)
	for _, v := range g.Nodes() {
		degrees[v] = len(g.AdjEdges(v))
	}

	adjMat := g.AsAdjMat()
	res := make([]float64, numNodes*numNodes)
	for i := 0; i < numNodes; i++ {
		for j := 0; j < numNodes; j++ {
			if i == j {
				res[i*numNodes+j] = float64(degrees[i])
			} else {
				res[i*numNodes+j] = -adjMat[i*numNodes+j]
			}
		}
	}

	return res
}
