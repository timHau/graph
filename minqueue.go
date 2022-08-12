package graph

import (
	"container/heap"
)

type Item struct {
	Node  int
	Prio  float64
	Index int
}

type MinQueue []*Item

func (mq MinQueue) Len() int { return len(mq) }

func (mq MinQueue) Less(i, j int) bool {
	return mq[i].Prio < mq[j].Prio
}

func (mq MinQueue) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
	mq[i].Index = i
	mq[j].Index = j
}

func (mq *MinQueue) Push(x any) {
	n := len(*mq)
	item := x.(*Item)
	item.Index = n
	*mq = append(*mq, item)
}

func (mq *MinQueue) Pop() any {
	old := *mq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*mq = old[0 : n-1]
	return item
}

func (mq *MinQueue) FindNode(node int) *Item {
	for _, item := range *mq {
		if item.Node == node {
			return item
		}
	}
	return nil
}

func (mq *MinQueue) UpdatePrio(item *Item, prio float64) {
	item.Prio = prio
	heap.Fix(mq, item.Index)
}
