package graph

import (
	"testing"
)

// Example Graph:
// .       7       ┌─────┐      9      ┌─────┐      1
// .  ┌───────────►│  1  ├────────────►│  3  ├──────────────┐
// .  │            └──┬──┘             └─────┘              │
// .  │               │                   ▲                 │
// .  │               │                   │                 │
// .  │               │                   │                 ▼
// ┌──┴──┐            │2                  │4             ┌─────┐
// │  0  │            │                   │              │  5  │
// └──┬──┘            │                   │              └─────┘
// .  │               │                   │                 ▲
// .  │               │                   │                 │
// .  │               ▼                   │                 │
// .  │    12      ┌─────┐      10     ┌──┴──┐      5       │
// .  └───────────►│  2  ├────────────►│  4  ├──────────────┘
// .               └─────┘             └─────┘
//
// Solution Graph:    7                  16
// .       7       ┌─────┐      9      ┌─────┐      1
// .  ┌───────────►│  1  ├────────────►│  3  ├──────────────┐
// .  │            └──┬──┘             └─────┘              │
// .  │               │                                     │
// .  │               │                                     │
// .  │               │                                     ▼
// ┌──┴──┐            │2                                 ┌─────┐
// │  0  │            │                                  │  5  │ 17
// └─────┘            │                                  └─────┘
// .                  │
// .                  │
// .                  ▼
// .               ┌─────┐      10     ┌─────┐
// .               │  2  ├────────────►│  4  │
// .               └─────┘             └─────┘
// .                  9                  19
func TestDijkstra(t *testing.T) {
	adjList := AdjList{
		0: []WeightTuple{{1, 7}, {2, 12}},
		1: []WeightTuple{{2, 2}, {3, 9}},
		2: []WeightTuple{{4, 10}},
		3: []WeightTuple{{5, 1}},
		4: []WeightTuple{{3, 4}, {5, 5}},
		5: []WeightTuple{},
	}
	g := FromAdjList(adjList)
	dist, pre, err := g.Dijkstra(0)
	if err != nil {
		t.Error(err)
	}

	expectDist := []float64{0, 7, 9, 16, 19, 17}
	expectPre := []int{-1, 0, 1, 1, 2, 3}

	for i := 0; i < len(dist); i++ {
		if dist[i] != expectDist[i] {
			t.Errorf("expected: %v, got: %v", expectDist, dist)
		}
		if pre[i] != expectPre[i] {
			t.Errorf("expected: %v, got: %v", expectPre, pre)
		}
	}
}

func TestDijkstraNegative(t *testing.T) {
	adjList := AdjList{
		0: []WeightTuple{{1, 7}, {2, 12}},
		1: []WeightTuple{{2, 2}, {3, 9}},
		2: []WeightTuple{{4, 10}},
		3: []WeightTuple{{5, -1}},
		4: []WeightTuple{{3, 4}, {5, 5}},
		5: []WeightTuple{},
	}
	g := FromAdjList(adjList)
	_, _, err := g.Dijkstra(0)
	if err == nil {
		t.Errorf("Dijkstra should not work with negative weights")
	}
}
