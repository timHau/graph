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
	mq = append(mq, &Item{Prio: 1})
	mq = append(mq, &Item{Prio: 2})
	heap.Init(&mq)
	if !mq.Less(0, 1) {
		t.Errorf("expected true, got false")
	}
}

func TestSwap(t *testing.T) {
	mq := MinQueue{}
	mq = append(mq, &Item{Prio: 0})
	mq = append(mq, &Item{Prio: 1})
	heap.Init(&mq)
	mq.Swap(0, 1)
	if mq[0].Prio != 1 || mq[1].Prio != 0 {
		t.Errorf("expected [1, 0], got %v", mq)
	}
}

func TestPush(t *testing.T) {
	mq := MinQueue{}
	mq.Push(&Item{Prio: 0})
	if mq[0].Prio != 0 {
		t.Errorf("expected [0], got %v", mq)
	}
}

func TestPop(t *testing.T) {
	mq := make(MinQueue, 0)
	mq.Push(&Item{Prio: 0, Node: 0, Index: 0})
	mq.Push(&Item{Prio: 1, Node: 1, Index: 1})
	mq.Push(&Item{Prio: 5, Node: 2, Index: 2})
	heap.Init(&mq)

	item := heap.Pop(&mq).(*Item)
	if item.Prio != 0 {
		t.Errorf("expected 0, got %f", item.Prio)
	}
	item = heap.Pop(&mq).(*Item)
	if item.Prio != 1 {
		t.Errorf("expected 0, got %f", item.Prio)
	}
	item = heap.Pop(&mq).(*Item)
	if item.Prio != 5 {
		t.Errorf("expected 0, got %f", item.Prio)
	}
}

func TestUpdatePrio(t *testing.T) {
	mq := make(MinQueue, 0)
	mq.Push(&Item{Prio: 0, Node: 0})
	mq.Push(&Item{Prio: 1, Node: 1})
	mq.Push(&Item{Prio: 5, Node: 2})
	heap.Init(&mq)

	item := mq.FindNode(2)
	mq.UpdatePrio(item, 10)
	if mq[2].Prio != 10 {
		t.Errorf("expected 2, got %f", mq[2].Prio)
	}
}
