package gograph

import "container/heap"

// Given an start node, find the shortest path to all other nodes in the graph
// Time Complexity: O(V^2)
// Space Complexity: O(V)
func (g *Graph[T, N]) Dijkstra(start int) {
	// Each entry in the distance array represents the distance from the start node to the node at the index
	distances := make([]int, len(g.Nodes))
	distances[start] = 0

	// List of predecessors for each node
	pre := make([]int, len(g.Nodes))
	pre[start] = start

	mq := make(MinQueue[T], len(g.Nodes))
	mq.Push(&Item[T]{
		node: g.Nodes[start],
		prio: distances[start],
	})

	// initialize the lists and the priority queue
	for i := 0; i < len(g.Nodes); i++ {
		if i != start {
			distances[i] = int(^uint(0) >> 1)
			pre[i] = -1
		}
		mq.Push(&Item[T]{
			node: g.Nodes[i],
			prio: distances[i],
		})
	}
	heap.Init(&mq)

	// while the priority queue is not empty
	for mq.Len() > 0 {
		// get the node with the smallest distance
		// item := heap.Pop(&mq).(*Item[T])

	}

}
