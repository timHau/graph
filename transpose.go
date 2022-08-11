package graph

func (g *Graph) Transpose() *Graph {
	tg := NewGraph()
	for _, edge := range g.Edges() {
		tg.AddEdge(edge.To, edge.From, edge.Weight)
	}

	return tg
}
