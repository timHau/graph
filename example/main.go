package main

import (
	"fmt"

	"github.com/timHau/gograph"
)

func printBfs(node int) {
	fmt.Printf("%d ", node)
}

func main() {
	g, err := gograph.FromAdjMat([]float64{
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
