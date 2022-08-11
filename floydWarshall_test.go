package graph

import (
	"math"
	"reflect"
	"testing"
)

// Example Graph
// ┌─────┐        10         ┌─────┐
// │  0  ├──────────────────►│  3  │
// └──┬──┘                   └─────┘
// .  │                         ▲
// .  │                         │
// . 5│                         │1
// .  │                         │
// .  ▼                         │
// ┌─────┐         3         ┌──┴──┐
// │  1  ├──────────────────►│  2  │
// └─────┘                   └─────┘
func TestFloydWarshall(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 3, 10)
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 1)

	m := g.FloydWarshall()
	expect := []float64{
		0, 5, 8, 9,
		math.Inf(1), 0, 3, 4,
		math.Inf(1), math.Inf(1), 0, 1,
		math.Inf(1), math.Inf(1), math.Inf(1), 0,
	}

	if !reflect.DeepEqual(m, expect) {
		t.Errorf("expected: %v, got: %v", expect, m)
	}
}

// Example Graph:
// .            ┌─────┐
// .  ┌────────►│  0  ├─────────┐
// .  │         └─────┘         │
// .  │                         │
// .  │4                      -2│
// .  │                         │
// .  │                         ▼
// ┌──┴──┐         3         ┌─────┐
// │  1  ├──────────────────►│  2  │
// └─────┘                   └──┬──┘
// .  ▲                         │
// .  │                         │
// .  │-1                      2│
// .  │                         │
// .  │         ┌─────┐         │
// .  └─────────┤  3  │◄────────┘
// .            └─────┘
func TestFloydWarshall2(t *testing.T) {
	g := FromEdgeList([]Edge{
		{1, 0, 4},
		{1, 2, 3},
		{0, 2, -2},
		{2, 3, 2},
		{3, 1, -1},
	})

	m := g.FloydWarshall()
	expect := []float64{
		0, -1, -2, 0,
		4, 0, 2, 4,
		5, 1, 0, 2,
		3, -1, 1, 0,
	}

	if !reflect.DeepEqual(m, expect) {
		t.Errorf("expected: %v, got: %v", expect, m)
	}
}
