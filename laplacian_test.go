package graph

import (
	"reflect"
	"testing"
)

// Example Graph
// .               ┌─────┐             ┌─────┐
// .  ┌────────────┤  3  ├─────────────┤  4  ├───────────┐
// .  │            └──┬──┘             └──┬──┘           │
// .  │               │                   │              │
// .  │               │                   │              │
// .  │               │                   │              │
// .  │               │                   │              │
// ┌──┴──┐            │                   │           ┌──┴──┐
// │  5  │            │                   │           │  0  │
// └─────┘            │                   │           └──┬──┘
// .                  │                   │              │
// .                  │                   │              │
// .                  │                   │              │
// .                  │                   │              │
// .               ┌──┴──┐             ┌──┴──┐           │
// .               │  2  ├─────────────┤  1  ├───────────┘
// .               └─────┘             └─────┘
func TestLaplacian(t *testing.T) {
	adjMat := []float64{
		0, 1, 0, 0, 1, 0,
		1, 0, 1, 0, 1, 0,
		0, 1, 0, 1, 0, 0,
		0, 0, 1, 0, 1, 1,
		1, 1, 0, 1, 0, 0,
		0, 0, 0, 1, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	laplacian := g.Laplacian()
	expected := []float64{
		2, -1, 0, 0, -1, 0,
		-1, 3, -1, 0, -1, 0,
		0, -1, 2, -1, 0, 0,
		0, 0, -1, 3, -1, -1,
		-1, -1, 0, -1, 3, 0,
		0, 0, 0, -1, 0, 1,
	}

	if !reflect.DeepEqual(laplacian, expected) {
		t.Errorf("expected laplacian = %v, got %v", expected, laplacian)
	}
}

// Example Graph
// .                 ┌─────┐
// .  ┌──────────────┤  0  ├───────────────┐
// .  │              └──┬──┘               │
// .  │                 │                  │
// .  │                 │                  │
// ┌──┴──┐              │               ┌──┴──┐
// │  1  │              │               │  2  │
// └─────┘              │               └──┬──┘
// .                    │                  │
// .                    │                  │
// .                 ┌──┴──┐               │
// .                 │  3  ├───────────────┘
// .                 └─────┘
func TestLaplacian2(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 1,
		1, 0, 0, 0,
		1, 0, 0, 1,
		1, 0, 1, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	laplacian := g.Laplacian()
	expected := []float64{
		3, -1, -1, -1,
		-1, 1, 0, 0,
		-1, 0, 2, -1,
		-1, 0, -1, 2,
	}

	if !reflect.DeepEqual(laplacian, expected) {
		t.Errorf("expected laplacian = %v, got %v", expected, laplacian)
	}
}
