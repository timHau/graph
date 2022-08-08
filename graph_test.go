package gograph

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g, err := NewGraph([]float64{
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
	_, err := NewGraph([]float64{
		1, 0, 1,
		0, 1, 0,
	}, []int{0, 1, 2})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestNewGraphFailSizeNodeVal(t *testing.T) {
	_, err := NewGraph([]float64{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	}, []int{0})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestAddNode(t *testing.T) {
	g := &Graph{}
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)

	if len(g.Nodes) != 3 {
		t.Errorf("expected 3 nodes, got %d", len(g.Nodes))
	}

	if g.Nodes[0].ID != 0 {
		t.Errorf("expected 0, got %d", g.Nodes[0].ID)
	}
	if g.Nodes[1].ID != 1 {
		t.Errorf("expected 1, got %d", g.Nodes[1].ID)
	}
	if g.Nodes[2].ID != 2 {
		t.Errorf("expected 2, got %d", g.Nodes[2].ID)
	}
}

func TestAddEdge(t *testing.T) {
	g := &Graph{}
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
	g := &Graph{}
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

func TestFromAdjList(t *testing.T) {
	AdjList := AdjList{
		0: []AdjListTuple{{1, 1}, {2, 1}},
		1: []AdjListTuple{{0, 3}},
		2: []AdjListTuple{{1, 2}},
	}
	g, err := FromAdjList(AdjList, []int{0, 1, 2})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(g.Nodes) != 3 {
		t.Errorf("expected 3 nodes, got %d", len(g.Nodes))
	}

	if len(g.Edges) != 4 {
		t.Errorf("expected 2 edges, got %d", len(g.Edges))
	}

	if g.Edges[0].From != 0 || g.Edges[0].To != 1 || g.Edges[0].Weight != 1 {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", g.Edges[0].From, g.Edges[0].Weight)
	}
	if g.Edges[1].From != 0 || g.Edges[1].To != 2 || g.Edges[1].Weight != 1 {
		t.Errorf("expected edge from 0 -- 2, got %d -- %.2f", g.Edges[1].From, g.Edges[1].Weight)
	}
	if g.Edges[2].From != 1 || g.Edges[2].To != 0 || g.Edges[2].Weight != 3 {
		t.Errorf("expected edge from 1 -- 0, got %d -- %.2f", g.Edges[2].From, g.Edges[2].Weight)
	}
	if g.Edges[3].From != 2 || g.Edges[3].To != 1 || g.Edges[3].Weight != 2 {
		t.Errorf("expected edge from 2 -- 1, got %d -- %.2f", g.Edges[3].From, g.Edges[3].Weight)
	}
}

func TestAdjList(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 1,
		1, 0, 0, 0,
		1, 0, 0, 0,
		1, 0, 0, 0,
	}
	g, err := NewGraph(adjMat, []int{1, 1, 1, 1})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	adjList := g.AsAdjList()
	expected := map[int][]AdjListTuple{
		0: {{1, 1}, {2, 1}, {3, 1}},
		1: {{0, 1}},
		2: {{0, 1}},
		3: {{0, 1}},
	}

	for k, v := range adjList {
		if !reflect.DeepEqual(v, expected[k]) {
			t.Errorf("expected adjList[%d] = %v, got %v", k, expected[k], v)
		}
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
func TestNeighbors(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, 1,
		1, 0, 0, 0,
		1, 0, 0, 0,
		1, 0, 0, 0,
	}
	g, err := NewGraph(adjMat, []int{1, 1, 1, 1})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	neighbors := g.Neighbors(0)
	expected := []AdjListTuple{
		{1, 1},
		{2, 1},
		{3, 1},
	}
	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("expected neighbors of 0 = %v, got %v", expected, neighbors)
	}
}

func TestEdge(t *testing.T) {
	adjMat := []float64{
		0, 3, 1, 1,
		1, 0, 0, 0,
		1, 0, 0, 0,
		1, 0, 0, 0,
	}
	g, err := NewGraph(adjMat, []int{1, 1, 1, 1})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	edge := g.Edge(0, 1)
	if edge.From != 0 || edge.To != 1 || edge.Weight != 3 {
		t.Errorf("expected edge from 0 ---> 1, got %d ---> %d", edge.From, edge.To)
	}
}
