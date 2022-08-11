package graph

import (
	"reflect"
	"sort"
	"testing"
)

// Example Graph:
// ┌─────┐         ┌─────┐                       ┌─────┐                   ┌─────┐
// │  0  │◄────────┤  1  │               ┌───────┤  5  │                   │  7  │
// └──┬──┘         └─────┘               │       └─────┘                   └──┬──┘
// .  │               ▲                  │            ▲                       │
// .  │               │                  │            │                       │
// .  │               │                  │            │                       │
// .  │               │                  │            └───────────┐           │
// .  │               │                  │                        │           │
// .  ▼               │                  ▼                        │           │
// ┌─────┐         ┌──┴──┐            ┌─────┐                  ┌──┴──┐        │
// │  3  ├────────►│  2  │◄───────────┤  4  ├─────────────────►│  6  │◄───────┘
// └─────┘         └─────┘            └─────┘                  └─────┘
func TestKosaraju2(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 0, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(4, 2, 1)
	g.AddEdge(4, 6, 1)
	g.AddEdge(5, 4, 1)
	g.AddEdge(6, 5, 1)
	g.AddEdge(7, 6, 1)

	scc := g.Kosaraju()
	if len(scc) != 3 {
		t.Errorf("expected 3 scc, got %d", len(scc))
	}
	sort.Ints(scc[0])
	sort.Ints(scc[1])
	sort.Ints(scc[2])

	sort.Slice(scc, func(i, j int) bool {
		if len(scc[i]) == 0 && len(scc[j]) == 0 {
			return false
		}
		if len(scc[i]) == 0 || len(scc[j]) == 0 {
			return len(scc[i]) == 0
		}

		return scc[i][0] < scc[j][0]
	})

	expect := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6},
		{7},
	}
	if !reflect.DeepEqual(scc, expect) {
		t.Errorf("expected %v, got %v", expect, scc)
	}
}
