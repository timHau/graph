package gograph

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g, err := FromAdjMat([]float64{
		1, 0, 1,
		0, 1, 0,
		1, 0, 1,
	})
	if err != nil {
		t.Error(err)
	}

	if g.NumNodes() != 3 {
		t.Errorf("expected 3 nodes, got %d", g.NumNodes())
	}

	if g.NumEdges() != 5 {
		t.Errorf("expected 5 edges, got %d", g.NumEdges())
	}
}

func TestNewGraphFailSizeAdj(t *testing.T) {
	_, err := FromAdjMat([]float64{
		1, 0, 1,
		0, 1, 0,
	})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestAddNode(t *testing.T) {
	g := New()
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)

	if g.NumNodes() != 3 {
		t.Errorf("expected 3 nodes, got %d", g.NumNodes())
	}
}

func TestAddEdge(t *testing.T) {
	g := New()
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)

	if g.NumEdges() != 3 {
		t.Errorf("expected 3 edges, got %d", g.NumEdges())
	}

	edge := g.Edge(0, 1)
	if edge.From != 0 || edge.To != 1 {
		t.Errorf("expected edge from 0 ---> 1, got %d ---> %d", edge.From, edge.To)
	}
	edge = g.Edge(1, 2)
	if edge.From != 1 || edge.To != 2 {
		t.Errorf("expected edge from 1 ---> 2, got %d ---> %d", edge.From, edge.To)
	}
	edge = g.Edge(2, 0)
	if edge.From != 2 || edge.To != 0 {
		t.Errorf("expected edge from 2 ---> 0, got %d ---> %d", edge.From, edge.To)
	}
}

func TestAddWeightedEdge(t *testing.T) {
	g := New()
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)
	g.AddWeightedEdge(0, 1, 2.1)
	g.AddWeightedEdge(1, 2, 1.8)
	g.AddWeightedEdge(2, 0, 9.1)

	if g.NumEdges() != 3 {
		t.Errorf("expected 3 edges, got %d", g.NumEdges())
	}

	edge := g.Edge(0, 1)
	if edge.From != 0 || edge.To != 1 || edge.Weight != 2.1 {
		t.Errorf("expected edge from 0 -- 2.1 --> 1, got %d -- %.2f --> %d", edge.From, edge.Weight, edge.To)
	}
	edge = g.Edge(1, 2)
	if edge.From != 1 || edge.To != 2 || edge.Weight != 1.8 {
		t.Errorf("expected edge from 1 -- 1.8 --> 2, got %d -- %.2f --> %d", edge.From, edge.Weight, edge.To)
	}
	edge = g.Edge(2, 0)
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
	g := FromAdjList(AdjList)

	if g.NumNodes() != 3 {
		t.Errorf("expected 3 nodes, got %d", g.NumNodes())
	}

	if g.NumEdges() != 4 {
		t.Errorf("expected 2 edges, got %d", g.NumEdges())
	}

	edge := g.Edge(0, 1)
	if edge.From != 0 || edge.To != 1 || edge.Weight != 1 {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", edge.From, edge.Weight)
	}
	edge = g.Edge(0, 2)
	if edge.From != 0 || edge.To != 2 || edge.Weight != 1 {
		t.Errorf("expected edge from 0 -- 2, got %d -- %.2f", edge.From, edge.Weight)
	}
	edge = g.Edge(1, 0)
	if edge.From != 1 || edge.To != 0 || edge.Weight != 3 {
		t.Errorf("expected edge from 1 -- 0, got %d -- %.2f", edge.From, edge.Weight)
	}
	edge = g.Edge(2, 1)
	if edge.From != 2 || edge.To != 1 || edge.Weight != 2 {
		t.Errorf("expected edge from 2 -- 1, got %d -- %.2f", edge.From, edge.Weight)
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
	g, err := FromAdjMat(adjMat)
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
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	edge := g.Edge(0, 1)
	if edge.From != 0 || edge.To != 1 || edge.Weight != 3 {
		t.Errorf("expected edge from 0 ---> 1, got %d ---> %d", edge.From, edge.To)
	}
}
