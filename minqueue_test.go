package graph

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
	heap.Init(&mq)
	if !mq.Less(0, 1) {
		t.Errorf("expected true, got false")
	}
}

func TestSwap(t *testing.T) {
	mq := MinQueue{}
	mq = append(mq, &Item{prio: 0})
	mq = append(mq, &Item{prio: 1})
	heap.Init(&mq)
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
	mq.Push(&Item{prio: 0, node: 0, index: 0})
	mq.Push(&Item{prio: 1, node: 1, index: 1})
	mq.Push(&Item{prio: 5, node: 2, index: 2})
	heap.Init(&mq)

	item := heap.Pop(&mq).(*Item)
	if item.prio != 0 {
		t.Errorf("expected 0, got %f", item.prio)
	}
	item = heap.Pop(&mq).(*Item)
	if item.prio != 1 {
		t.Errorf("expected 0, got %f", item.prio)
	}
	item = heap.Pop(&mq).(*Item)
	if item.prio != 5 {
		t.Errorf("expected 0, got %f", item.prio)
	}
}

func TestUpdatePrio(t *testing.T) {
	mq := make(MinQueue, 0)
	mq.Push(&Item{prio: 0, node: 0})
	mq.Push(&Item{prio: 1, node: 1})
	mq.Push(&Item{prio: 5, node: 2})
	heap.Init(&mq)

	item := mq.FindNode(2)
	mq.UpdatePrio(item, 10)
	if mq[2].prio != 10 {
		t.Errorf("expected 2, got %f", mq[2].prio)
	}
}
