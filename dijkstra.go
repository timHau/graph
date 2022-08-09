package gograph

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
// 2)  a Graph containing the shortest paths to all other nodes as edge weights
func (g *Graph) Dijkstra(start int) (*Graph, error) {
	if g.HasNegativeEdges() {
		return nil, errors.New("Dijkstras Algorithm does not support negative edge weights")
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
		prio:  distances[start],
		node:  start,
		index: start,
	})

	// initialize the lists and the priority queue
	for i := range nodes {
		if i != start {
			distances[i] = math.MaxFloat64
			pre[i] = -1
			mq.Push(&Item{
				prio:  distances[i],
				node:  i,
				index: i,
			})
		}
	}
	heap.Init(&mq)

	// while the priority queue is not empty
	for mq.Len() > 0 {
		// get the node with the smallest distance
		item := heap.Pop(&mq).(*Item)
		u := item.node
		for _, e := range g.AdjEdges(u) {
			v := e.To
			vItem := mq.FindNode(v)
			if vItem != nil {
				alt := distances[u] + e.Weight
				if alt < distances[v] && distances[u] != math.MaxFloat64 {
					distances[v] = alt
					pre[v] = item.node
					mq.UpdatePrio(vItem, alt)
				}
			}
		}
	}

	res := NewGraph()
	for i, v := range pre {
		if v != -1 {
			res.AddWeightedEdge(v, i, distances[i])
		}
	}

	return res, nil
}
