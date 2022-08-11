package graph

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

func TestFromEdgeList(t *testing.T) {
	edgeList := []Edge{
		{0, 1, 1},
		{1, 2, 1},
		{2, 0, 1},
	}
	g := FromEdgeList(edgeList)
	if g.NumNodes() != 3 {
		t.Errorf("expected 3 nodes, got %d", g.NumNodes())
	}

	edge := g.Edge(0, 1)
	if !edge.Equals(&edgeList[0]) {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", g.Edge(0, 1).From, g.Edge(0, 1).Weight)
	}
	edge = g.Edge(1, 2)
	if !edge.Equals(&edgeList[1]) {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", g.Edge(0, 1).From, g.Edge(0, 1).Weight)
	}
	edge = g.Edge(2, 0)
	if !edge.Equals(&edgeList[2]) {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", g.Edge(0, 1).From, g.Edge(0, 1).Weight)
	}
}

func TestAddNode(t *testing.T) {
	g := NewGraph()
	g.AddNode(0)
	g.AddNode(1)
	g.AddNode(2)

	if g.NumNodes() != 3 {
		t.Errorf("expected 3 nodes, got %d", g.NumNodes())
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 0, 1)

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

func TestAddEdge2(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1, 2.1)
	g.AddEdge(1, 2, 1.8)
	g.AddEdge(2, 0, 9.1)

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
		0: []WeightTuple{{1, 1}, {2, 1}},
		1: []WeightTuple{{0, 3}},
		2: []WeightTuple{{1, 2}},
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
// .│  0  ├────────────┤  1  │
// .└──┬──┤            └─────┘
// .   │  └─────┐
// .   │        │
// .   │        │
// .   │        │
// .   │        │
// .┌──┴──┐     │      ┌─────┐
// .│  3  │     └──────┤  2  │
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

	neighbors := g.AdjEdges(0)
	expected := []Edge{
		{0, 1, 1},
		{0, 2, 1},
		{0, 3, 1},
	}
	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("expected neighbors of 0 = %v, got %v", expected, neighbors)
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
func TestNeighbors2(t *testing.T) {
	adjMat := []float64{
		0, 1, 0,
		0, 0, 1,
		0, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	e0 := g.AdjEdges(0)
	if len(e0) != 1 {
		t.Errorf("expected 1 neighbor for node 0, got %d", len(e0))
	}
	if e0[0].From != 0 || e0[0].To != 1 || e0[0].Weight != 1 {
		t.Errorf("expected edge from 0 -- 1, got %d -- %.2f", e0[0].From, e0[0].Weight)
	}

	e1 := g.AdjEdges(1)
	if len(e1) != 1 {
		t.Errorf("expected 1 neighbors for node 1, got %d", len(e1))
	}
	if e1[0].From != 1 || e1[0].To != 2 || e1[0].Weight != 1 {
		t.Errorf("expected edge from 1 -- 2, got %d -- %.2f", e1[0].From, e1[0].Weight)
	}

	e2 := g.AdjEdges(2)
	if len(e2) != 0 {
		t.Errorf("expected 0 neighbors for node 2, got %d", len(e2))
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

func TestFromZeroAdjMat(t *testing.T) {
	adjMat := []float64{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if g.NumNodes() != 4 {
		t.Errorf("expected 4 nodes, got %d", g.NumNodes())
	}
}

func TestHasNoNegativeEdges(t *testing.T) {
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

	if g.HasNegativeEdges() {
		t.Errorf("expected no negative edges, got %v", g.HasNegativeEdges())
	}
}

func TestHasNegativeEdges(t *testing.T) {
	adjMat := []float64{
		0, 1, 1, -1,
		1, 0, 0, 0,
		1, 0, 0, 0,
		1, 0, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !g.HasNegativeEdges() {
		t.Errorf("expected negative edges, got %v", g.HasNegativeEdges())
	}
}
