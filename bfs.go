package graph

// Breadth First Search
//
// Traverse the graph while exploring all nodes at the same level
// only explore nodes that have not been visited
//
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *Graph) BFS(start int, fn func(int)) {
	numNodes := g.NumNodes()
	visited := make([]bool, numNodes)
	g.BFSstep(start, visited, fn)

	// If the graph is not connected, we will start exploring the remaining graph components
	for i := 0; i < numNodes; i++ {
		if !visited[i] {
			g.BFSstep(i, visited, fn)
		}
	}
}

// Breadth First Step
//
// Perfomes a breadth first search step starting at node start
func (g *Graph) BFSstep(start int, visited []bool, fn func(int)) {
	queue := make([]int, 0)
	queue = append(queue, start)
	visited[start] = true
	fn(start)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, edge := range g.Edges() {
			if edge.From == curr && !visited[edge.To] {
				queue = append(queue, edge.To)
				visited[edge.To] = true
				fn(edge.To)
			}
		}
	}
}
