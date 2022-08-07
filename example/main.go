package main

import (
	"fmt"

	"github.com/timHau/gograph"
)

func printBfs(node *gograph.Node[int], _ int) {
	fmt.Printf("%d ", node.Val)
}

func main() {
	g, err := gograph.NewGraph(6, []int{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	}, []int{1, 2, 3, 4, 5, 6})
	if err != nil {
		panic(err)
	}

	g.BreadthFirstSearch(0, printBfs)
}