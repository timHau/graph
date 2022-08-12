package main

import (
	"fmt"

	"github.com/timHau/graph"
)

func main() {
	edgeList := []graph.Edge{
		{0, 1, 1},
		{1, 2, 1},
		{2, 0, 1},
	}
	gL := graph.FromEdgeList(edgeList)
	fmt.Println("Graph:", gL)
	fmt.Println()

	g2 := graph.NewGraph()
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 0, 1)
	fmt.Println("Graph:", g2)
	fmt.Println()

	g3 := graph.FromAdjList(graph.AdjList{
		0: []graph.WeightTuple{{1, 1}, {2, 1}},
		1: []graph.WeightTuple{{0, 3}},
		2: []graph.WeightTuple{{1, 2}},
	})
	fmt.Println("Graph:", g3)
	fmt.Println()

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

	fmt.Println("Laplacian:")
	n := g.NumNodes()
	laplacian := g.Laplacian()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%.2f ", laplacian[i*n+j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Breadth first search, starting at node 0:")
	g.BFS(0, func(i int) {
		if i != 5 {
			fmt.Print(i, " -> ")
		} else {
			fmt.Print(i)
		}
	})
	fmt.Println()
	fmt.Println()

	fmt.Println("Depth first search, starting at node 0:")
	g.DFS(0, func(i int) {
		if i != 5 {
			fmt.Print(i, " -> ")
		} else {
			fmt.Print(i)
		}
	})
	fmt.Println()
	fmt.Println()

	fmt.Println("All pairs shortest path:")
	dist := g.FloydWarshall()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%.2f ", dist[i*n+j])
		}
		fmt.Println()
	}
	fmt.Println()
}
