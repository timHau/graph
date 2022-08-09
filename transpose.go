package gograph

func (g *Graph) Transpose() *Graph {
	numNodes := g.NumNodes()
	adjMat := make([]float64, numNodes*numNodes)
	for i := 0; i < numNodes; i++ {
		for _, edge := range g.AdjEdges(i) {
			adjMat[edge.To*numNodes+i] = edge.Weight
		}
	}
	// we guarantee that the matrix is square so no need to handle the err
	tg, _ := FromAdjMat(adjMat)
	return tg
}
