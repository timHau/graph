package main

import (
	"fmt"

	"github.com/timHau/go-graph/graph"
)

func main() {
	g, err := graph.NewGraph(3, []int{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	}, []int{0, 1, 2})
	if err != nil {
		panic(err)
	}

	fmt.Println(g)
}
