package gograph

import (
	"container/heap"
	"testing"
)

func TestLen(t *testing.T) {
	mq := MinQueue{}
	if mq.Len() != 0 {
		t.Errorf("expected 0, got %d", mq.Len())
	}
}

func TestLess(t *testing.T) {
	mq := MinQueue{}
	mq = append(mq, &Item{prio: 1})
	mq = append(mq, &Item{prio: 2})
	if mq.Less(0, 1) {
		t.Errorf("expected true, got false")
	}
}

func TestSwap(t *testing.T) {
	mq := MinQueue{}
	mq = append(mq, &Item{prio: 0})
	mq = append(mq, &Item{prio: 1})
	mq.Swap(0, 1)
	if mq[0].prio != 1 || mq[1].prio != 0 {
		t.Errorf("expected [1, 0], got %v", mq)
	}
}

func TestPush(t *testing.T) {
	mq := MinQueue{}
	mq.Push(&Item{prio: 0})
	if mq[0].prio != 0 {
		t.Errorf("expected [0], got %v", mq)
	}
}

func TestPop(t *testing.T) {
	mq := make(MinQueue, 0)
	mq.Push(&Item{prio: 0})
	mq.Push(&Item{prio: 1})
	mq.Push(&Item{prio: 5})
	heap.Init(&mq)

	item := mq.Pop().(*Item)
	if item.prio != 0 {
		t.Errorf("expected 0, got %f", item.prio)
	}
}

func TestUpdatePrio(t *testing.T) {
	mq := make(MinQueue, 0)
	mq.Push(&Item{prio: 0})
	mq.Push(&Item{prio: 1})
	mq.Push(&Item{prio: 5})
	heap.Init(&mq)

	mq.UpdatePrio(mq[0], 2)
	if mq[0].prio != 2 {
		t.Errorf("expected 2, got %f", mq[0].prio)
	}
}
