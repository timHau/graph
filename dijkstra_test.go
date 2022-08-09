package gograph

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
	d, pre := g.Dijkstra(0)

	expectedD := []float64{0, 7, 9, 16, 19, 17}
	expectedPre := []int{0, 0, 1, 1, 2, 3}

	for i := 0; i < len(d); i++ {
		if d[i] != expectedD[i] {
			t.Errorf("expected distance %f, got %f", expectedD[i], d[i])
		}
		if pre[i] != expectedPre[i] {
			t.Errorf("expected predecessor %d, got %d", expectedPre[i], pre[i])
		}
	}
}
