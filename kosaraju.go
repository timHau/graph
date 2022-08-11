package graph

// Kosaraju's algorithm for finding strongly connected components.
//
// returns a list of strongly connected components (scc), each of which is a list of nodes
//
// Time Complexity: O(V + E)
func (g *Graph) Kosaraju() [][]int {
	s := make([]int, 0)
	visited := make([]bool, g.NumNodes())

	for _, v := range g.Nodes() {
		if !visited[v] {
			g.DFSstep(v, visited, func(n int) {
				s = append(s, n)
			})
		}
	}

	tg := g.Transpose()

	scc := make([][]int, 0)
	visited = make([]bool, g.NumNodes())

	for len(s) > 0 {
		// pop from the stack
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if !visited[v] {
			scc = append(scc, make([]int, 0))
			// explore the component
			tg.DFSstep(v, visited, func(n int) {
				scc[len(scc)-1] = append(scc[len(scc)-1], n)
			})
		}
	}

	return scc
}
