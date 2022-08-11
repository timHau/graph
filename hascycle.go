package graph

func (g *Graph) HasCycle() bool {
	numNodes := g.NumNodes()
	visited := make([]bool, numNodes)
	recStack := make([]bool, numNodes)

	for i := 0; i < numNodes; i++ {
		if !visited[i] {
			if g.detectCycle(i, visited, recStack) {
				return true
			}
		}
	}

	return false
}

func (g *Graph) detectCycle(start int, visited, recStack []bool) bool {
	visited[start] = true
	recStack[start] = true
	for _, edge := range g.AdjEdges(start) {
		if !visited[edge.To] {
			if g.detectCycle(edge.To, visited, recStack) {
				return true
			}
		} else if recStack[edge.To] {
			return true
		}
	}
	recStack[start] = false
	return false
}
