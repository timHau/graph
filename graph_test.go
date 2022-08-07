package gograph

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g, err := NewGraph(3, []int{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	}, []int{0, 1, 2})
	if err != nil {
		t.Error(err)
	}

	if len(g.Nodes) != 3 {
		t.Errorf("expected 3 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 5 {
		t.Errorf("expected 5 edges, got %d", len(g.Edges))
	}
}

func TestNewGraphFailSizeAdj(t *testing.T) {
	_, err := NewGraph(3, []int{
		1, 0, 1,
		0, 1, 0,
	}, []int{0, 1, 2})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewGraphFailSizeNodeVal(t *testing.T) {
	_, err := NewGraph(3, []int{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	}, []int{0})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestAddNode(t *testing.T) {
	g := &Graph[int, int]{}
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)

	if len(g.Nodes) != 3 {
		t.Errorf("expected 3 nodes, got %d", len(g.Nodes))
	}

	if g.Nodes[0].Val != 0 {
		t.Errorf("expected 0, got %d", g.Nodes[0].Val)
	}
	if g.Nodes[1].Val != 1 {
		t.Errorf("expected 1, got %d", g.Nodes[1].Val)
	}
	if g.Nodes[2].Val != 2 {
		t.Errorf("expected 2, got %d", g.Nodes[2].Val)
	}
}

func TestAddEdge(t *testing.T) {
	g := &Graph[int, int]{}
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)

	if len(g.Edges) != 3 {
		t.Errorf("expected 3 edges, got %d", len(g.Edges))
	}
	if g.Edges[0].From != 0 || g.Edges[0].To != 1 {
		t.Errorf("expected edge from 0 ---> 1, got %d ---> %d", g.Edges[0].From, g.Edges[0].To)
	}
	if g.Edges[1].From != 1 || g.Edges[1].To != 2 {
		t.Errorf("expected edge from 1 ---> 2, got %d ---> %d", g.Edges[1].From, g.Edges[1].To)
	}
	if g.Edges[2].From != 2 || g.Edges[2].To != 0 {
		t.Errorf("expected edge from 2 ---> 0, got %d ---> %d", g.Edges[2].From, g.Edges[2].To)
	}
}

func TestAddWeightedEdge(t *testing.T) {
	g := &Graph[int, float32]{}
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddWeightedEdge(0, 1, 2.1)
	g.AddWeightedEdge(1, 2, 1.8)
	g.AddWeightedEdge(2, 0, 9.1)

	if len(g.Edges) != 3 {
		t.Errorf("expected 3 edges, got %d", len(g.Edges))
	}

	edge := g.Edges[0]
	if edge.From != 0 || edge.To != 1 || edge.Weight != 2.1 {
		t.Errorf("expected edge from 0 -- 2.1 --> 1, got %d -- %.2f --> %d", edge.From, edge.Weight, edge.To)
	}
	edge = g.Edges[1]
	if edge.From != 1 || edge.To != 2 || edge.Weight != 1.8 {
		t.Errorf("expected edge from 1 -- 1.8 --> 2, got %d -- %.2f --> %d", edge.From, edge.Weight, edge.To)
	}
	edge = g.Edges[2]
	if edge.From != 2 || edge.To != 0 || edge.Weight != 9.1 {
		t.Errorf("expected edge from 2 -- 9.1 --> 0, got %d -- %.2f --> %d", edge.From, edge.Weight, edge.To)
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
func TestBFS(t *testing.T) {
	adjMat := []int{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	}
	g, _ := NewGraph(6, adjMat, []int{1, 2, 3, 4, 5, 6})

	res := make([]int, 0)
	g.BreadthFirstSearch(0, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})

	if reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) == false {
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
	adjMat := []int{
		0, 1, 1, 0, 0, 0,
		1, 0, 0, 1, 1, 0,
		1, 0, 0, 0, 1, 0,
		0, 1, 0, 0, 1, 1,
		0, 1, 1, 1, 0, 1,
		0, 0, 0, 1, 1, 0,
	}
	g, _ := NewGraph(6, adjMat, []int{1, 2, 3, 4, 5, 6})

	res := make([]int, 0)
	g.BreadthFirstSearch(1, func(n *Node[int], _ int) {
		res = append(res, n.Val)
	})

	if reflect.DeepEqual(res, []int{2, 1, 4, 5, 3, 6}) == false {
		t.Errorf("expected [1, 2, 3, 4, 5, 6], got %v", res)
	}
}

// Example Graph:
// .┌─────┐            ┌─────┐
// .│  1  ├────────────┤  2  │
// .└──┬──┤            └─────┘
// .   │  └─────┐
// .   │        │
// .   │        │
// .   │        │
// .   │        │
// .┌──┴──┐     │      ┌─────┐
// .│  4  │     └──────┤  3  │
// .└─────┘            └─────┘
func TestArbitraryNodeTypes(t *testing.T) {
	type Coordinate struct {
		X, Y int
	}

	adjMat := []int{
		0, 1, 1, 1,
		1, 0, 0, 0,
		1, 0, 0, 0,
		1, 0, 0, 0,
	}
	nodeVals := []Coordinate{
		{0, 0}, {1, 0}, {0, 1}, {1, 1},
	}
	g, err := NewGraph(4, adjMat, nodeVals)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(g.Nodes) != 4 {
		t.Errorf("expected 4 nodes, got %d", len(g.Nodes))
	}

	for i, nv := range nodeVals {
		node := g.Nodes[i]
		if node.Val != nv {
			t.Errorf("expected node %v, got %v", nv, node.Val)
		}
	}
}
