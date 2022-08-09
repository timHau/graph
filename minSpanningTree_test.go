package gograph

import "testing"

func TestPrim(t *testing.T) {
	adjMat := []float64{
		0, 4, 0, 0, 0, 0, 0, 8, 0,
		4, 0, 8, 0, 0, 0, 0, 11, 0,
		0, 8, 0, 7, 0, 4, 0, 0, 2,
		0, 0, 7, 0, 9, 14, 0, 0, 0,
		0, 0, 0, 9, 0, 10, 0, 0, 0,
		0, 0, 4, 14, 10, 0, 2, 0, 0,
		0, 0, 0, 0, 0, 2, 0, 1, 6,
		8, 11, 0, 0, 0, 0, 1, 0, 7,
		0, 0, 2, 0, 0, 0, 6, 7, 0,
	}
	g, err := FromAdjMat(adjMat)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	mst := g.Prim()
	t.Errorf("%v", mst)
}
