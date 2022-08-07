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
	adjMat := []int{
		0, 1, 1, 1, 0,
		1, 0, 1, 0, 0,
		1, 1, 0, 0, 1,
		1, 0, 0, 0, 0,
		0, 0, 1, 0, 0,
	}
	nodeVals := []int{0, 1, 2, 3, 4}
	g, err := NewGraph(adjMat, nodeVals)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res := make([]int, 0)
	g.DFS(0, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})
	if !reflect.DeepEqual(res, []int{0, 1, 2, 4, 3}) {
		t.Errorf("expected [0, 1, 2, 4, 3], got %v", res)
	}
}