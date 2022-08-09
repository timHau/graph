package gograph

import (
	"reflect"
	"testing"
)

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
func TestTranspose(t *testing.T) {
	adjMat := []float64{
		0, 1, 0,
		0, 0, 1,
		1, 0, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	tg := g.Transpose()
	expectedAdjList := AdjList{
		0: []WeightTuple{{2, 1}},
		1: []WeightTuple{{0, 1}},
		2: []WeightTuple{{1, 1}},
	}
	if !reflect.DeepEqual(tg.AdjacencyList, expectedAdjList) {
		t.Errorf("expected %v, got %v", expectedAdjList, tg.AdjacencyList)
	}
}

func TestTransposeTrivial(t *testing.T) {
	adjMat := []float64{1}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	tg := g.Transpose()
	expectedAdjList := AdjList{
		0: []WeightTuple{{0, 1}},
	}
	if !reflect.DeepEqual(tg.AdjacencyList, expectedAdjList) {
		t.Errorf("expected %v, got %v", expectedAdjList, tg.AdjacencyList)
	}
}
