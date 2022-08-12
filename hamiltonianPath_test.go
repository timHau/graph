package graph

import "testing"

// Example Graph
// .           ┌─────┐
// .           │  0  │
// .           └──┬──┘
// .              │
// .              │
// .              │
// .           ┌──┴──┐
// .  ┌────────┤  1  ├──────────┐
// .  │        └─────┘          │
// .  │                         │
// .  │                         │
// .  │                         │
// ┌──┴──┐                   ┌──┴──┐
// │  2  ├───────────────────┤  3  │
// └─────┘                   └─────┘
func TestHamiltonianPath(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 0, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, 1)
	g.AddEdge(3, 1, 1)

	hasHP := g.HasHamiltonianPathDP()
	if !hasHP {
		t.Errorf("expected hasHamiltonianPath = true, got false")
	}
}

// Example Graph
// .           ┌─────┐
// .           │  0  │
// .           └──┬──┘
// .              │
// .              │
// .              │
// .           ┌──┴──┐
// .  ┌────────┤  1  ├──────────┐
// .  │        └─────┘          │
// .  │                         │
// .  │                         │
// .  │                         │
// ┌──┴──┐                   ┌──┴──┐
// │  2  │                   │  3  │
// └─────┘                   └─────┘
func TestNoHamiltonianPath(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 0, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(3, 1, 1)

	hasHP := g.HasHamiltonianPathDP()
	if hasHP {
		t.Errorf("expected hasHamiltonianPath = false, got true")
	}
}
