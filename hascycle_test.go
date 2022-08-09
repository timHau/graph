package gograph

import "testing"

// Example Graph:
// ┌──────┐                       ┌──────┐
// │      │                       │      │
// │   ┌──┴──┐     ┌─────┐     ┌──┴──┐   │
// └──►│  0  │     │  1  │     │  2  │◄──┘
// .   └─────┘     └─────┘     └─────┘
func TestHasCycle(t *testing.T) {
	adjMat := []float64{
		1, 0, 0,
		0, 0, 0,
		0, 0, 1,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !g.HasCycle() {
		t.Errorf("expected true, got false")
	}
}

// Example Graph:
// .           ┌─────┐
// .  ┌───────►│  1  ├────────┐
// .  │        └─────┘        │
// .  │                       │
// .  │                       │
// .  │                       ▼
// ┌──┴──┐                 ┌─────┐
// │  0  │◄────────────────┤  2  │
// └─────┘                 └─────┘
func TestHasCycle2(t *testing.T) {
	adjMat := []float64{
		0, 1, 0,
		0, 0, 1,
		1, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !g.HasCycle() {
		t.Errorf("expected true, got false")
	}
}

// Example Graph:
// ┌─────┐                 ┌─────┐
// │  1  ├────────────────▲│  2  ├───┐
// └─────┘                ◄└──┬──┘   │
// .  ▲       ┌───────────┘   │      │
// .  │       │               │      │
// .  │       │               │      │
// .  │       │               │      │
// .  │       │               ▼      │
// ┌──┴──┐    │            ┌─────┐   │
// │  0  ├────┘        ┌───┤  3  │   │
// └─────┘             │   └─────┘   │
// .  ▲                │      ▲      │
// .  │                └──────┘      │
// .  │                              │
// .  └──────────────────────────────┘
func TestHasCycle3(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0,
		0, 0, 1, 0,
		1, 0, 0, 1,
		0, 0, 0, 1,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !g.HasCycle() {
		t.Errorf("expected true, got false")
	}
}

// Example Graph:
// .          ┌─────┐
// .  ┌──────►│  1  ├─────────┐
// .  │       └─────┘         │
// .  │                       │
// .  │                       │
// .  │                       │
// .  │                       │
// .  │                       ▼
// ┌──┴──┐                 ┌─────┐
// │  0  │                 │  2  │
// └─────┘                 └─────┘
func TestHasNoCycle(t *testing.T) {
	adjMat := []float64{
		0, 1, 0,
		0, 0, 1,
		0, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if g.HasCycle() {
		t.Errorf("expected false, got true")
	}
}

// Example Graph:
// ┌─────┐                 ┌─────┐
// │  1  ├────────────────►│  2  │
// └─────┘                 └──┬──┘
// .  ▲                    ▲  │
// .  │                    │  │
// .  │          ┌─────────┘  │
// .  │          │            │
// .  │          │            │
// .  │          │            │
// .  │          │            ▼
// ┌──┴──┐       │         ┌─────┐
// │  0  ├───────┘         │  3  │
// └─────┘                 └─────┘
func TestHasNoCycle2(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
		0, 0, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if g.HasCycle() {
		t.Errorf("expected false, got true")
	}
}
