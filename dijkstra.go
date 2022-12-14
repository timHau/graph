package graph

import (
	"container/heap"
	"errors"
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
func (g *Graph) Dijkstra(start int) ([]float64, []int, error) {
	if g.HasNegativeEdges() {
		return nil, nil, errors.New("dijkstras Algorithm does not support negative edge weights")
	}

	nodes := g.Nodes()
	numNodes := len(nodes)
	// Each entry in the distance array represents the distance from the start node to the node at the index
	distances := make([]float64, numNodes)
	distances[start] = 0

	// List of predecessors for each node
	pre := make([]int, numNodes)
	pre[start] = -1

	mq := make(MinQueue, 0)
	mq.Push(&Item{
		Prio:  distances[start],
		Node:  start,
		Index: start,
	})

	// initialize the lists and the priority queue
	for i := range nodes {
		if i != start {
			distances[i] = math.MaxFloat64
			pre[i] = -1
			mq.Push(&Item{
				Prio:  distances[i],
				Node:  i,
				Index: i,
			})
		}
	}
	heap.Init(&mq)

	// while the priority queue is not empty
	for mq.Len() > 0 {
		// get the node with the smallest distance
		item := heap.Pop(&mq).(*Item)
		u := item.Node
		for _, e := range g.AdjEdges(u) {
			v := e.To
			vItem := mq.FindNode(v)
			if vItem != nil {
				alt := distances[u] + e.Weight
				if alt < distances[v] && distances[u] != math.MaxFloat64 {
					distances[v] = alt
					pre[v] = item.Node
					mq.UpdatePrio(vItem, alt)
				}
			}
		}
	}

	return distances, pre, nil
}
