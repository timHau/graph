package graph

import (
	"reflect"
	"testing"
)

// Example Graph:
// ┌─────┐        -1           ┌─────┐        2
// │  0  ├────────────────────►│  1  ├───────────────┐
// └──┬──┘                     ├────┬┘               │
// .  │                        │ ▲  │                │
// .  │           3            │ │  │                ▼
// . 4│ ┌──────────────────────┘ │  │             ┌─────┐
// .  │ │                        │  │             │  4  │
// .  │ │                       1│  │2            └──┬──┘
// .  │ │                        │  │                │
// .  ▼ ▼                        │  ▼                │
// ┌─────┐           5         ┌─┴───┐       -3      │
// │  2  │◄────────────────────┤  3  │◄──────────────┘
// └─────┘                     └─────┘
func TestBellmanFord(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, -1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 2)
	g.AddEdge(1, 4, 2)
	g.AddEdge(3, 2, 5)
	g.AddEdge(3, 1, 1)
	g.AddEdge(4, 3, -3)

	dist, err := g.BellmanFord(0)
	if err != nil {
		t.Error(err)
	}

	expect := []float64{0, -1, 2, -2, 1}
	if !reflect.DeepEqual(dist, expect) {
		t.Errorf("expected: %v, got: %v", expect, dist)
	}
}
