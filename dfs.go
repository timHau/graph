package graph

// # Depth First Search
//
//	Traverse the graph while following the edges in a depth-first manner
//
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) DFS(start int, fn func(int)) {
	numNodes := g.NumNodes()
	visited := make([]bool, numNodes)
	g.DFSstep(start, visited, fn)

	// If the graph is not connected, we will start exploring the remaining graph components
	for i := 0; i < numNodes; i++ {
		if !visited[i] {
			g.DFSstep(i, visited, fn)
		}
	}
}

func (g *Graph) DFSstep(start int, visited []bool, fn func(int)) {
	visited[start] = true
	fn(start)
	for _, edge := range g.AdjEdges(start) {
		if !visited[edge.To] {
			g.DFSstep(edge.To, visited, fn)
		}
	}
}
