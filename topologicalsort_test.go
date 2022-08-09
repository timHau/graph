package gograph

import (
	"reflect"
	"testing"
)

// Example Graph:
//
// .           ┌─────┐                 ┌─────┐
// .  ┌────────┤  5  ├──────┐   ┌──────┤  4  ├──────┐
// .  │        └─────┘      │   │      └─────┘      │
// .  │                     │   │                   │
// .  │                     │   │                   │
// .  │                     │   │                   │
// .  │                     │   │                   │
// .  │                     │   │                   │
// .  ▼                     ▼   ▼                   ▼
// ┌─────┐                 ┌─────┐               ┌─────┐
// │  2  │                 │  0  │               │  1  │
// └──┬──┘                 └─────┘               └─────┘
// .  │                                             ▲
// .  │                                             │
// .  │                                             │
// .  │                    ┌─────┐                  │
// .  └───────────────────►│  3  ├──────────────────┘
// .                       └─────┘
func TestTopologicalSort(t *testing.T) {
	adjList := AdjList{
		0: []WeightTuple{},
		1: []WeightTuple{},
		2: []WeightTuple{{3, 1}},
		3: []WeightTuple{{1, 1}},
		4: []WeightTuple{{0, 1}, {1, 1}},
		5: []WeightTuple{{0, 1}, {2, 1}},
	}
	g := FromAdjList(adjList)
	ts, err := g.TopologicalSort()
	if err != nil {
		t.Errorf("TopologicalSort() error: %v", err)
	}

	expect := []int{5, 4, 2, 3, 1, 0}
	if !reflect.DeepEqual(ts, expect) {
		t.Errorf("TopologicalSort() = %v, want %v", ts, expect)
	}
}

// Example Graph:
// .             ┌─────┐                       ┌─────┐
// .  ┌─────────►│  1  ├──────────────────────►│  4  │
// .  │          └──┬──┘                       └─────┘
// .  │             │                             ▲
// .  │             └───────────┐  ┌──────────────┘
// .  │                         ▼  │
// ┌──┴──┐                     ┌───┴─┐
// │  0  │                     │  3  │
// └──┬──┘                     └───┬─┘
// .  │                         ▲  │
// .  │             ┌───────────┘  └──────────────┐
// .  │             │                             ▼
// .  │          ┌──┴──┐                       ┌─────┐
// .  └─────────►│  2  ├──────────────────────►│  5  │
// .             └─────┘                       └─────┘
func TestTopologicalSort2(t *testing.T) {
	adjList := AdjList{
		0: []WeightTuple{{1, 1}, {2, 1}},
		1: []WeightTuple{{3, 1}, {4, 1}},
		2: []WeightTuple{{3, 1}, {5, 1}},
		3: []WeightTuple{{4, 1}, {5, 1}},
		4: []WeightTuple{},
		5: []WeightTuple{},
	}
	g := FromAdjList(adjList)
	ts, err := g.TopologicalSort()
	if err != nil {
		t.Errorf("TopologicalSort() error: %v", err)
	}

	// there are multiple valid topological sorts
	expect := []int{0, 2, 1, 3, 5, 4}
	if !reflect.DeepEqual(ts, expect) {
		t.Errorf("TopologicalSort() = %v, want %v", ts, expect)
	}
}
