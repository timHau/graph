package gograph

import (
	"container/heap"
	"math"
)

// Dijkstra's algorithm for single source shortest paths
//
// Given an start node, find the shortest path to all other nodes in the graph
// Time Complexity: O(V^2)
// Space Complexity: O(V)
// returns:
// 1)  a list of shortest distances to all other nodes
// 2)  a list of predecessors for each node
func (g *Graph) Dijkstra(start int) ([]float64, []int) {
	numNodes := g.NumNodes()
	// Each entry in the distance array represents the distance from the start node to the node at the index
	distances := make([]float64, numNodes)
	distances[start] = 0

	// List of predecessors for each node
	pre := make([]int, numNodes)
	pre[start] = start

	mq := make(MinQueue, numNodes)
	mq.Push(&Item{
		prio:  distances[start],
		index: start,
	})

	// initialize the lists and the priority queue
	for i := 0; i < numNodes; i++ {
		if i != start {
			distances[i] = math.MaxFloat64
			pre[i] = -1
		}
		mq.Push(&Item{
			prio:  distances[i],
			index: i,
		})
	}
	heap.Init(&mq)

	// while the priority queue is not empty
	for mq.Len() > 0 {
		// get the node with the smallest distance
		item := heap.Pop(&mq).(*Item)
		for n := range g.Neighbors(item.index) {
			alt := distances[item.index] + float64(g.Edge(item.index, n).Weight)
			if alt < distances[n] && distances[item.index] != math.MaxFloat64 {
				distances[n] = alt
				pre[n] = item.index
				mq.UpdatePrio(item, alt)
			}
		}
	}

	return distances, pre
}
