package gograph

import (
	"reflect"
	"testing"
)

// Example Graph:
// ┌─────┐              ┌─────┐
// │  0  ├──────────────┤  3  │
// └──┬─┬┘              └─────┘
// .  │ │
// .  │ │
// .  │ └────┐
// .  │      ├─────┐
// .  │ ┌────┤  2  │
// .  │ │    └──┬──┘
// .  │ │       │
// .  │ │       │
// .  │ │       │
// ┌──┴─┴┐      │       ┌─────┐
// │  1  │      └───────┤  4  │
// └─────┘              └─────┘
func TestDFS(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 1, 0,
		1, 0, 1, 0, 0,
		1, 1, 0, 0, 1,
		1, 0, 0, 0, 0,
		0, 0, 1, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.DFS(0, func(n int) {
		res = append(res, n)
	})
	if !reflect.DeepEqual(res, []int{0, 1, 2, 4, 3}) {
		t.Errorf("expected [0, 1, 2, 4, 3], got %v", res)
	}
}

// Example Graph:
// .         ┌─────┐
// .  ┌──────┤  0  ├───────┐
// .  │      └─────┘       │
// .  │                    │
// .  │                    │
// ┌──┴──┐              ┌──┴──┐
// │  2  ├──────────────┤  1  │
// └──┬─┬┘              └┬─┬──┘
// .  │ │                │ │
// .  │ │                │ │
// .  │ └────┬─────┬─────┘ │
// .  │      │  3  │       │
// .  │ ┌────┴─────┴─────┐ │
// .  │ │                │ │
// .  │ │                │ │
// ┌──┴─┴┐              ┌┴─┴──┐
// │  4  │              │  5  │
// └─────┘              └─────┘
func TestDFS2(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0, 0, 0,
		1, 0, 1, 1, 0, 1,
		1, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 1, 1,
		0, 0, 1, 1, 0, 0,
		0, 1, 0, 1, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.DFS(0, func(n int) {
		res = append(res, n)
	})

	if !reflect.DeepEqual(res, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("expected [0, 1, 2, 3, 4, 5], got %v", res)
	}
}

// Example Graph:
// ┌─────┐        ┌─────┐
// │  0  │    ┌───┤  1  │
// └──┬──┘    │   └──┬──┘
// .  │       │      │
// .  │       └──────┘
// .  │                   ┌─────┐
// .  │                   │     │
// ┌──┴──┐             ┌──┴──┐  │
// │  2  ├─────────────┤  3  ├──┘
// └─────┘             └─────┘
func TestDFSDisconnected(t *testing.T) {
	adjMat := []float64{
		0, 0, 1, 0,
		0, 1, 0, 0,
		1, 0, 0, 1,
		0, 0, 1, 1,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.DFS(0, func(n int) {
		res = append(res, n)
	})
	if !reflect.DeepEqual(res, []int{0, 2, 3, 1}) {
		t.Errorf("expected [0, 2, 3, 1], got %v", res)
	}
}
