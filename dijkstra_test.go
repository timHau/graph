package gograph

import (
	"reflect"
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
	g = g.Dijkstra(0)

	expected := AdjList{
		0: []WeightTuple{{1, 7}},
		1: []WeightTuple{{2, 9}, {3, 16}},
		2: []WeightTuple{{4, 19}},
		3: []WeightTuple{{5, 17}},
	}
	if !reflect.DeepEqual(g.AdjacencyList, expected) {
		t.Errorf("expected %v, got %v", expected, g.AdjacencyList)
	}
}
