package graph

import "errors"

func (g *Graph) TopologicalSort() ([]int, error) {
	if g.HasCycle() {
		return nil, errors.New("graph has cycle")
	}
	visited := make([]bool, g.NumNodes())
	stack := make([]int, 0)

	for i := 0; i < g.NumNodes(); i++ {
		if !visited[i] {
			g.TopologicalStep(i, visited, &stack)
		}
	}

	// reverse the stack
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack, nil
}

func (g *Graph) TopologicalStep(node int, visited []bool, stack *[]int) {
	visited[node] = true
	for _, e := range g.AdjEdges(node) {
		if !visited[e.To] {
			g.TopologicalStep(e.To, visited, stack)
		}
	}
	*stack = append(*stack, node)
}
