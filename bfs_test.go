package gograph

import (
	"reflect"
	"testing"
)

// Example Graph:
// .          ┌─────┐
// .  ┌───────┤  1  ├───────┐
// .  │       └─────┘       │
// .  │                     │
// .  │                     │
// ┌──┴──┐               ┌──┴──┐
// │  2  │               │  3  │
// └──┬──┴──────┐        └──┬──┘
// .  │         │           │
// .  │         └────────┐  │
// ┌──┴──┐               ├──┴──┐
// │  4  ├───────────────┤  5  │
// └──┬──┘               └──┬──┘
// .  │                     │
// .  │                     │
// .  │       ┌─────┐       │
// .  └───────┤  6  ├───────┘
// .          └─────┘
func TestBFS(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	}
	g, err := NewGraph(adjMat, []int{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	res := make([]int, 0)
	g.BFS(0, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("expected [1, 2, 3, 4, 5, 6], got %v", res)
	}
}

// Example Graph:
// .          ┌─────┐
// .  ┌───────┤  1  ├───────┐
// .  │       └─────┘       │
// .  │                     │
// .  │                     │
// ┌──┴──┐               ┌──┴──┐
// │  2  │               │  3  │
// └──┬──┴──────┐        └──┬──┘
// .  │         │           │
// .  │         └────────┐  │
// ┌──┴──┐               ├──┴──┐
// │  4  ├───────────────┤  5  │
// └──┬──┘               └──┬──┘
// .  │                     │
// .  │                     │
// .  │       ┌─────┐       │
// .  └───────┤  6  ├───────┘
// .          └─────┘
func TestBFS2(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	}
	g, err := NewGraph(adjMat, []int{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.BFS(1, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})

	if !reflect.DeepEqual(res, []int{2, 1, 4, 5, 3, 6}) {
		t.Errorf("expected [1, 2, 3, 4, 5, 6], got %v", res)
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
func TestBFS3(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 0, 0, 0,
		1, 0, 1, 1, 0, 1,
		1, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 1, 1,
		0, 0, 1, 1, 0, 0,
		0, 1, 0, 1, 0, 0,
	}
	nodeVals := []int{0, 1, 2, 3, 4, 5}
	g, err := NewGraph(adjMat, nodeVals)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.BFS(0, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})

	if !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 4}) {
		t.Errorf("expected [0, 1, 2, 3, 5, 4], got %v", res)
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
func TestBFSDisconnected(t *testing.T) {
	adjMat := []float64{
		0, 0, 1, 0,
		0, 1, 0, 0,
		1, 0, 0, 1,
		0, 0, 1, 1,
	}
	g, err := NewGraph(adjMat, []int{0, 1, 2, 3})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.BFS(0, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})
	if !reflect.DeepEqual(res, []int{0, 2, 3, 1}) {
		t.Errorf("expected [0, 2, 3, 1], got %v", res)
	}
}
