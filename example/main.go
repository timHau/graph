package main

import (
	"fmt"

	"github.com/timHau/graph"
)

func printBfs(node int) {
	fmt.Printf("%d ", node)
}

func main() {
	edgeList := []graph.Edge{
		{0, 1, 1},
		{1, 2, 1},
		{2, 0, 1},
	}
	gL := graph.FromEdgeList(edgeList)
	fmt.Println("Graph:", gL)

	g2 := graph.NewGraph()
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 0, 1)
	fmt.Println("Graph:", g2)

	g3 := graph.FromAdjList(graph.AdjList{
		0: []graph.WeightTuple{{1, 1}, {2, 1}},
		1: []graph.WeightTuple{{0, 3}},
		2: []graph.WeightTuple{{1, 2}},
	})
	fmt.Println("Graph:", g3)

	g, err := graph.FromAdjMat([]float64{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	})
	if err != nil {
		panic(err)
	}

	g.BFS(0, printBfs)
	fmt.Println()
}
