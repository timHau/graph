package graph

// Checks if the Graph has a Hamiltonian Path (a Path that visits every vertex exactly once)
// using dynamic programming with time complexity O((2^n)*n^2).
// Warning: dont use on large graphs (this problem is NP-complete)
func (g *Graph) HasHamiltonianPathDP() bool {
	// adapted from: https://www.hackerearth.com/practice/algorithms/graphs/hamiltonian-path/tutorial/
	// initialize the dp matrix. dp[j][i] checks if there is a path that visits each vertex in
	// the subset represented by the mask i and ends at vertex j.
	numNodes := g.NumNodes()
	numSubsets := 1 << uint(numNodes) // aka 2^numNodes
	dp := make([][]bool, numNodes)
	for i := 0; i < numNodes; i++ {
		dp[i] = make([]bool, numSubsets)
	}

	for i := 0; i < numNodes; i++ {
		// dp[i][2^i] represents a subset that consists of only vertex i. Clearly, there is a path
		dp[i][1<<uint(i)] = true // dp[i][2^i] = true
	}

	for i := 0; i < numSubsets; i++ { // for all subsets
		for j := 0; j < numNodes; j++ { // for all vertices
			if checkIthBit(j, i) { // check if vertex j is in subset i
				for k := 0; k < numNodes; k++ { // for all vertices
					// ccheck if vertex k is in subset i and is adjacent to vertex j
					// i^(1<<uint(j)) aka (i XOR 2^j) represents the subset S/{j}
					// the cell dp[k][i^(1<<uint(j))] represents whether there is a path
					// that visits each vertex in S/{j} exactly once and ends at vertex k.
					if k != j && checkIthBit(k, i) && g.Edge(k, j) != nil && dp[k][i^(1<<uint(j))] {
						dp[j][i] = true
						break
					}
				}
			}
		}
	}

	for i := 0; i < numNodes; i++ {
		if dp[i][(1<<uint(numNodes))-1] {
			return true
		}
	}

	return false
}

func checkIthBit(i int, mask int) bool {
	return (mask & (1 << uint(i))) != 0
}
