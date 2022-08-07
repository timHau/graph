package gograph

// Depth First Search
//
// # Traverse the graph while following the edges in a depth-first manner
//
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph[T]) DFS(start int, fn func(*Node[T], int)) {
	visited := make([]bool, len(g.Nodes))
	g.DFSstep(start, visited, fn)

	// If the graph is not connected, we will start exploring the remaining graph components
	for i := 0; i < len(g.Nodes); i++ {
		if !visited[i] {
			g.DFSstep(i, visited, fn)
		}
	}
}

func (g *Graph[T]) DFSstep(start int, visited []bool, fn func(*Node[T], int)) {
	visited[start] = true
	fn(g.Nodes[start], start)
	for _, edge := range g.Edges {
		if edge.From == start && !visited[edge.To] {
			g.DFSstep(edge.To, visited, fn)
		}
	}
}
