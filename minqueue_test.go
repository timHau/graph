package gograph

import (
	"container/heap"
	"testing"
)

func TestLen(t *testing.T) {
	mq := MinQueue[int]{}
	if mq.Len() != 0 {
		t.Errorf("expected 0, got %d", mq.Len())
	}
}

func TestLess(t *testing.T) {
	mq := MinQueue[int]{}
	mq = append(mq, &Item[int]{prio: 1})
	mq = append(mq, &Item[int]{prio: 2})
	if mq.Less(0, 1) {
		t.Errorf("expected true, got false")
	}
}

func TestSwap(t *testing.T) {
	mq := MinQueue[int]{}
	mq = append(mq, &Item[int]{prio: 0})
	mq = append(mq, &Item[int]{prio: 1})
	mq.Swap(0, 1)
	if mq[0].prio != 1 || mq[1].prio != 0 {
		t.Errorf("expected [1, 0], got %v", mq)
	}
}

func TestPush(t *testing.T) {
	mq := MinQueue[int]{}
	mq.Push(&Item[int]{prio: 0})
	if mq[0].prio != 0 {
		t.Errorf("expected [0], got %v", mq)
	}
}

func TestPop(t *testing.T) {
	mq := make(MinQueue[int], 0)
	mq.Push(&Item[int]{prio: 0})
	mq.Push(&Item[int]{prio: 1})
	mq.Push(&Item[int]{prio: 5})
	heap.Init(&mq)

	item := mq.Pop().(*Item[int])
	if item.prio != 0 {
		t.Errorf("expected 0, got %d", item.prio)
	}
}

func TestUpdate(t *testing.T) {
	mq := make(MinQueue[int], 0)
	mq.Push(&Item[int]{prio: 0})
	mq.Push(&Item[int]{prio: 1})
	mq.Push(&Item[int]{prio: 5})
	heap.Init(&mq)

	mq.update(mq[0], &Node[int]{}, 2)
	if mq[0].prio != 2 {
		t.Errorf("expected 2, got %d", mq[0].prio)
	}
}
